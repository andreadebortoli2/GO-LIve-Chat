package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/andreadebortoli2/GO-Live-Chat/internal/config"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/handlers"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/models"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/context"
	"github.com/gorilla/websocket"
	"github.com/justinas/nosurf"
)

func Router(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	// MIDDLEWARES
	// recover from panics, log the panic and return an HTTP 500 status(if possible)
	r.Use(middleware.Recoverer)
	// add crsf token in cookies
	r.Use(NoSurf)
	// load session
	r.Use(LoadSession)

	// public routes
	r.Get("/", handlers.Repo.ShowHomePage)
	r.Get("/login", handlers.Repo.ShowLoginPage)
	r.Post("/login", handlers.Repo.PostLogin)
	r.Get("/logout", handlers.Repo.ShowLogoutPage)
	r.Get("/new-user", handlers.Repo.ShowNewUserPage)
	r.Post("/new-user", handlers.Repo.PostNewUser)

	// secure routes
	r.Group(func(r chi.Router) {
		// auth middleware
		r.Use(authMiddleware)

		r.Get("/dashboard", handlers.Repo.ShowDashboardPage)
		r.Get("/profile", handlers.Repo.ShowProfilePage)
		r.Get("/chat", handlers.Repo.ShowChatPage)
		r.Get("/older-messages", handlers.Repo.ShowOlderMessages)
		// r.Post("/new-message", handlers.Repo.PostNewMessage) removed to use websocket
		r.Get("/ws", handlers.Repo.WebsocketHandler)

		r.Group(func(r chi.Router) {
			r.Use(moderatrorAuthMiddleware)
			r.Get("/moderators-list", handlers.Repo.ShowModeratorsPage)
		})

		// restricted admin routes
		r.Route("/admin", func(r chi.Router) {
			// users CRUD
			r.Use(adminAuthMiddleware)
			r.Get("/all-users", handlers.Repo.ShowAdminAllUsersPage)
			r.Post("/change-access-level", handlers.Repo.PostChangeAccessLevel)
			r.Post("/delete-user", handlers.Repo.PostDeleteUser)
		})
	})

	return r
}

// MIDDLEWARES
// NoSurf generate CRSF token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// LoadSession load the session
func LoadSession(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !websocket.IsWebSocketUpgrade(r) {

			ses, err := session.Get(r, "active_user")
			if err != nil {
				log.Println("cannot get the session (load session middleware)", err)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			context.Set(r, "active_user", ses)
			next.ServeHTTP(w, r)
			err = ses.Save(r, w)
			if err != nil {
				log.Println("cannot save session (post new user)", err)
			}

		} else {

			next.ServeHTTP(w, r)
		}
	})

}

// authMiddleware authenticate registered users
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ses, err := session.Get(r, "active_user")
		if err != nil {
			log.Println("cannot get the session (auth middleware)", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if ses.IsNew {
			// log.Println("not authorized (auth middleware)")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// adminAuthMiddleware give access to only admin section
func adminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ses, err := session.Get(r, "active_user")
		if err != nil {
			log.Println("cannot get the session (admin auth middleware)", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		u := ses.Values["user"].(models.User)
		if u.AccessLevel == "3" {
			next.ServeHTTP(w, r)
		} else {
			log.Println("not authorized (admin auth middleware)")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	})
}

// moderatrorAuthMiddleware give access to mod adn admin to moderator section
func moderatrorAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ses, err := session.Get(r, "active_user")
		if err != nil {
			log.Println("cannot get the session (moderator auth middleware)", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		u := ses.Values["user"].(models.User)
		accLvl, err := strconv.Atoi(u.AccessLevel)
		if err != nil {
			log.Println("cannot convert access levele (moderator auth middleware)")
		}
		if accLvl > 1 {
			next.ServeHTTP(w, r)
		} else {
			log.Println("not authorized (moderator auth middleware)")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	})
}
