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
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *TemplatesHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// renderiza login comum (se quiser, pode ler userType para ajustar textos)
	tmpl := template.Must(template.ParseFiles("../docs/loginpage/login.html"))
	_ = tmpl.Execute(w, nil)
}

func (h *TemplatesHandler) HandleRegistry(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, sessionName)
	userType, _ := sess.Values[userTypeKey].(string)
	if userType != "organizer" {
		userType = "common"
	}
	tmpl := template.Must(template.ParseFiles("../docs/register/register.html"))
	_ = tmpl.Execute(w, model.RegisterPageData{UserType: userType})
}
