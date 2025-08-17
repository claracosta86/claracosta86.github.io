package handlers

import (
	"net/http"
	"encoding/json"

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
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(model.Type{UserType: userType})
}

// user/get-type [GET]
func (h *TemplatesHandler) HandleGetUserType(w http.ResponseWriter, r *http.Request) {
    sess, _ := session.Store.Get(r, sessionName)
    userType, ok := sess.Values[userTypeKey].(string)
    if !ok || userType != "organizer" {
        userType = "common"
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(model.Type{UserType: userType})
}

// /user/set-information [POST]
func (h *TemplatesHandler) HandleSetUserInformation(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	userType := r.Form.Get("userType")
	if userType != "organizer" {
		userType = "common"
	}

	userID := r.Form.Get("userID")

	sess, _ := session.Store.Get(r, sessionName)
	sess.Values[userTypeKey] = userType
	sess.Values["userID"] = userID

	_ = sess.Save(r, w)
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(model.Information{UserType: userType, UserID: userID})
}

// user/get-type [GET]
func (h *TemplatesHandler) HandleGetUserInformation(w http.ResponseWriter, r *http.Request) {
    sess, _ := session.Store.Get(r, sessionName)
    userType, ok := sess.Values[userTypeKey].(string)
    if !ok || userType != "organizer" {
        userType = "common"
    }
	
	userID, ok := sess.Values["userID"].(string)
	if !ok {
		userID = ""
	}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(model.Information{UserType: userType, UserID: userID})
}