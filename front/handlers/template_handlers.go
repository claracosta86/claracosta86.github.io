package handlers

import (
	"html/template"
	"net/http"

	"poc2/front/session"
	"poc2/front/model"

)

const sessionName = "app-session"
const userTypeKey = "userType"

type TemplatesHandler struct {}

func NewTemplatesHandler() *TemplatesHandler {
	return &TemplatesHandler{}
}

// /users/select-type [POST]
func (h *TemplatesHandler) HandleUserTypeSelection(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	userType := r.Form.Get("userType")
	if userType != "organizer" {
		userType = "common"
	}

	sess, _ := session.Store.Get(r, sessionName)
	sess.Values[userTypeKey] = userType
	_ = sess.Save(r, w)
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

// /users/login [GET]
func (h *TemplatesHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, sessionName)
	userType, _ := sess.Values[userTypeKey].(string)
	if userType != "organizer" {
		userType = "common"
	}

	userID, _ := sess.Values["userID"].(string)
	if userID == "" {
		userID = "0" 
	}

	tmpl := template.Must(template.ParseFiles("./docs/loginpage/login.html"))
	_ = tmpl.Execute(w, model.LoginPageData{UserID: userID, UserType: userType})
}

// /users/register [GET]
func (h *TemplatesHandler) HandleRegistry(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, sessionName)
	userType, _ := sess.Values[userTypeKey].(string)
	if userType != "organizer" {
		userType = "common"
	}

	tmpl := template.Must(template.ParseFiles("./docs/registerpage/register.html"))
	_ = tmpl.Execute(w, model.RegisterPageData{UserType: userType})
}

// /users/password-recovery [GET]
func (h *TemplatesHandler) HandleForgottenPassword(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./docs/loginpage/password-recovery.html"))
	_ = tmpl.Execute(w, nil)
}

// /users/profile [GET]
func (h *TemplatesHandler) HandleProfile(w http.ResponseWriter, r *http.Request){
	sess, _ := session.Store.Get(r, sessionName)
	userType, _ := sess.Values[userTypeKey].(string)
	if userType != "organizer" {
		userType = "common"
	}

	userID, _ := sess.Values["userID"].(string)
	if userID == "" {
		userID = "0" 
	}
	tmpl := template.Must(template.ParseFiles("./docs/homepage/users/profile.html"))
	_ = tmpl.Execute(w, model.LoginPageData{UserID: userID, UserType: userType})
}

// /users/profile/edit [GET]
func (h *TemplatesHandler) HandleEditProfile(w http.ResponseWriter, r *http.Request)	{
	sess, _ := session.Store.Get(r, sessionName)
	userType, _ := sess.Values[userTypeKey].(string)
	if userType != "organizer" {
		userType = "common"
	}

	tmpl := template.Must(template.ParseFiles("./docs/homepage/users/edit-profile.html"))
	_ = tmpl.Execute(w, model.RegisterPageData{UserType: userType})
}

func (h *TemplatesHandler) HandleChangePassword(w http.ResponseWriter, r *http.Request)	{
	sess, _ := session.Store.Get(r, sessionName)
	userType, _ := sess.Values[userTypeKey].(string)
	if userType != "organizer" {
		userType = "common"
	}

	tmpl := template.Must(template.ParseFiles("./docs/homepage/users/change-password.html"))
	_ = tmpl.Execute(w, model.RegisterPageData{UserType: userType})
}

func (h *TemplatesHandler) HandleViewFavorites(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, sessionName)
	userType, _ := sess.Values[userTypeKey].(string)
	if userType != "organizer" {
		userType = "common"
	}

	tmpl := template.Must(template.ParseFiles("./docs/homepage/users/favorites.html"))
	_ = tmpl.Execute(w, model.RegisterPageData{UserType: userType})
}

// /home [GET]
func (h *TemplatesHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, sessionName)
	userType, _ := sess.Values[userTypeKey].(string)
	if userType != "organizer" {
		userType = "common"
	}

	userID, _ := sess.Values["userID"].(string)
	
	tmpl := template.Must(template.ParseFiles("./docs/homepage/home.html"))
	_ = tmpl.Execute(w, model.LoginPageData{UserID: userID, UserType: userType})
}