package main

import (
	"net/http"
	"strconv"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/handlers"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/models"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/justinas/nosurf"
)

func Router() *chi.Mux {
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
	r.Get("/about", handlers.Repo.ShowAboutPage)
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

		// restricted admin routes
		r.Route("/admin", func(r chi.Router) {
			r.Use(adminAuthMiddleware)
			r.Get("/all-users", handlers.Repo.ShowAdminAllUsersPage)
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
	return session.LoadAndSave(next)
}

// authMiddleware authenticate registered users
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userExist := session.Exists(r.Context(), "user")
		if !userExist {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		user, ok := session.Get(r.Context(), "user").(models.User)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		accLvl, err := strconv.Atoi(user.AccessLevel)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if accLvl > 0 {
			next.ServeHTTP(w, r)
		}
	})
}

// adminAuthMiddleware authenticate registered users
func adminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userExist := session.Exists(r.Context(), "user")
		if !userExist {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		user, ok := session.Get(r.Context(), "user").(models.User)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		accLvl, err := strconv.Atoi(user.AccessLevel)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if accLvl == 3 {
			next.ServeHTTP(w, r)
		}
	})
}
