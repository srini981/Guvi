package controllers

import (
	"encoding/json"
	"guviTask/api/services"
	"net/http"
)

func UserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method!="GET"{
		customErr:="methods not found"
		http.Error(w, customErr, http.StatusNotFound)
		return
	}

	email:=r.Header.Get("usermail")
	users,err:=services.RegistrationsSvc.CheckIfUserRegistered(email)
	if err != nil {
	customErr := "user not found"
	http.Error(w, customErr, http.StatusNotFound)
	return
  }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
