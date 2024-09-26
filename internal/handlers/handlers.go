package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/andreadebortoli2/GO-Live-Chat/internal/config"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/database"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/helpers"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/models"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/render"
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
	render.RenderPage(w, r, "home", render.TemplateData{})
}

// ShowLoginPage show login page
func (m *Repository) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "login", render.TemplateData{})
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
		helpers.RenderErr(err, w, r, "login", fields)
		return
	}

	user, err := database.Login(email, password)
	if err != nil {
		helpers.RenderErr(err, w, r, "login", fields)
		return
	}

	_ = m.App.Session.RenewToken(r.Context())

	m.App.Session.Put(r.Context(), "user", user)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// ShowLogoutPage logic to logout the user
func (m *Repository) ShowLogoutPage(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ShowNewUserPage show new-user page
func (m *Repository) ShowNewUserPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "new-user", render.TemplateData{})
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
		helpers.RenderErr(err, w, r, "new-user", fields)
		return
	}

	err = database.AddUser(userName, email, password)
	if err != nil {
		helpers.RenderErr(err, w, r, "new-user", fields)
		return
	}

	// after correct registration immediatly login the user and redirect
	user, err := database.Login(email, password)
	if err != nil {
		helpers.RenderErr(err, w, r, "login", fields)
		return
	}

	_ = m.App.Session.RenewToken(r.Context())

	m.App.Session.Put(r.Context(), "user", user)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// ShowDashboardPage show dashboard page
func (m *Repository) ShowDashboardPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "dashboard", render.TemplateData{})
}

// ShowProfilePage show user profile page
func (m *Repository) ShowProfilePage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "profile", render.TemplateData{})
}

// ShowAdminAllUsersPage show the administration page with all the users
func (m *Repository) ShowAdminAllUsersPage(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers()
	if err != nil {
		helpers.RenderErr(err, w, r, "admin-all-users", nil)
	}
	datausers := make(map[string]interface{})
	for i, u := range users {
		index := strconv.Itoa(i)
		datausers[index] = u
	}
	data := make(map[string]interface{})
	data["users"] = datausers
	render.RenderPage(w, r, "admin-all-users", render.TemplateData{
		Data: data,
	})
}

// PostChangeAccessLevel modify the access level of a selected user
func (m *Repository) PostChangeAccessLevel(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/admin/all-users", http.StatusSeeOther)
		return
	}

	accLvl := r.Form.Get("moderator")
	userID := r.Form.Get("user-id")

	err = database.SetModerator(accLvl, userID)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/admin/all-users", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/admin/all-users", http.StatusSeeOther)
}

// PostDeleteUser delete a selected user
func (m *Repository) PostDeleteUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/admin/all-users", http.StatusSeeOther)
		return
	}

	userID := r.Form.Get("user-id")

	database.DeleteUserByID(userID)

	http.Redirect(w, r, "/admin/all-users", http.StatusSeeOther)
}

// ShowChatPage show the chat page with last messages
func (m *Repository) ShowChatPage(w http.ResponseWriter, r *http.Request) {

	messages, err := database.GetLastMessages()
	if err != nil {
		helpers.RenderErr(err, w, r, "chat", nil)
		return
	}

	dataMessages := make(map[string]interface{})
	for i, m := range messages {
		index := strconv.Itoa(i)
		// add 0 to single digit index for a correct display on page
		if len(index) == 1 {
			index = "0" + index
		}
		dataMessages[index] = m
	}
	data := make(map[string]interface{})
	data["messages"] = dataMessages

	render.RenderPage(w, r, "chat", render.TemplateData{
		Data: data,
	})
}

func (m *Repository) ShowOlderMessages(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/chat", http.StatusSeeOther)
		return
	}

	messages, err := database.GetOlderMessages()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/chat", http.StatusSeeOther)
		return
	}

	var msgstempl string

	for _, m := range messages {
		u := Repo.App.Session.Get(r.Context(), "user").(models.User)
		activeUserID := u.ID
		var msgsHTMLstr string
		if m.User.ID == activeUserID {
			msgsHTMLstr = fmt.Sprintf(
				`<div class="d-flex justify-content-end">
				<div class="card w-75 mb-3 text-end bg-warning-subtle">
					<div class="card-body">
						<h6 class="card-title">%s</h6>
						<p class="card-text">%s</p>
					</div>
				</div>
			</div>
		`, m.User.UserName, m.Content)
		} else {
			msgsHTMLstr = fmt.Sprintf(
				`<div class="d-flex justify-content-start">
				<div class="card w-75 mb-3 text-start bg-success-subtle">
					<div class="card-body">
						<h6 class="card-title">%s</h6>
						<p class="card-text">%s</p>
					</div>
				</div>
			</div>
		`, m.User.UserName, m.Content)
		}

		msgstempl += msgsHTMLstr
	}

	// return with HTMX
	templ, _ := template.New("t").Parse(msgstempl)
	templ.Execute(w, nil)
}

func (m *Repository) PostNewMessage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/chat", http.StatusSeeOther)
		return
	}

	userId := r.Form.Get("user-id")
	msg := r.Form.Get("message-content")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/chat", http.StatusSeeOther)
		return
	}

	err = database.PostNewMessage(userIdInt, msg)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/chat", http.StatusSeeOther)
		return
	}

	u := Repo.App.Session.Get(r.Context(), "user").(models.User)
	msgStr := fmt.Sprintf(
		`<div class="d-flex justify-content-end">
				<div class="card w-75 mb-3 text-end bg-warning-subtle">
					<div class="card-body">
						<h6 class="card-title">%s</h6>
						<p class="card-text">%s</p>
					</div>
				</div>
			</div>
		`, u.UserName, msg)

	// return with HTMX
	templ, _ := template.New("t").Parse(msgStr)
	templ.Execute(w, nil)
}
