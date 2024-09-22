package handlers

import (
	"log"
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/config"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/database"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/helpers"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/models"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

// ShowHomePage show home page
func (m *Repository) ShowHomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "home", nil)
}

// ShowAboutPage show about page
func (m *Repository) ShowAboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "about", nil)
}

// ShowLoginPage show login page
func (m *Repository) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "login", nil)
}

// PostLogin logic to login the user
func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	fields := map[string]string{
		"email":    email,
		"password": password,
	}

	err = helpers.LoginValidation(fields)
	if err != nil {
		log.Println(err)
		fields["error"] = err.Error()
		render.RenderPage(w, r, "login", fields)
		return
	}

	user, err := database.Login(email, password)
	if err != nil {
		log.Println(err)
		fields["error"] = err.Error()
		render.RenderPage(w, r, "login", fields)
		return
	}

	// if a user is returned give the authorization saving auth level in the session
	if user != (models.User{}) {
		_ = m.App.Session.RenewToken(r.Context())

		m.App.Session.Put(r.Context(), "userName", user.UserName)
		m.App.Session.Put(r.Context(), "accessLevel", user.AccessLevel)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// ShowLogoutPage logic to logout the user
func (m *Repository) ShowLogoutPage(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ShowNewUserPage show new-user page
func (m *Repository) ShowNewUserPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "new-user", nil)
}

// PostNewUserPage add new user to DB
func (m *Repository) PostNewUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/new-user", http.StatusSeeOther)
		return
	}

	userName := r.Form.Get("user_name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	fields := map[string]string{
		"user_name": userName,
		"email":     email,
		"password":  password,
	}

	err = helpers.NewUserValidation(fields)
	if err != nil {
		log.Println(err)
		fields["error"] = err.Error()
		render.RenderPage(w, r, "new-user", fields)
		return
	}

	err = database.AddUser(userName, email, password)
	if err != nil {
		log.Println(err)
		fields["error"] = err.Error()
		render.RenderPage(w, r, "new-user", fields)
		return
	}

	// after correct registration immediatly login the user and redirect
	user, err := database.Login(email, password)
	if err != nil {
		log.Println(err)
		fields["error"] = err.Error()
		render.RenderPage(w, r, "new-user", fields)
		return
	}
	if user != (models.User{}) {
		_ = m.App.Session.RenewToken(r.Context())
		m.App.Session.Put(r.Context(), "auth", user.AccessLevel)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// ShowDashboardPage show dashboard page
func (m *Repository) ShowDashboardPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "dashboard", nil)
}

// ShowAdminAllUsersPage show the administraation page with all the users
func (m *Repository) ShowAdminAllUsersPage(w http.ResponseWriter, r *http.Request) {
	// TODO: display the users
	render.RenderPage(w, r, "admin-all-users", nil)
}
