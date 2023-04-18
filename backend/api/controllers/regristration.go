package controllers

import (
	"encoding/json"
	"guviTask/api/controllers/models"
	"guviTask/api/services"
	"io"

	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request){
	if r.Method!="POST"{
		customErr:="methods not found"
		http.Error(w, customErr, http.StatusNotFound)
		return
	}

	user:=models.Userprofile{}

	arr,err:=io.ReadAll(r.Body)
	if err != nil{
		customErr:="failed to parse request body"
		http.Error(w, customErr, http.StatusBadRequest)
		return
	}

	err=json.Unmarshal(arr,&user)

	if err != nil{
		customErr:="failed to parse request body"
		http.Error(w, customErr, http.StatusBadRequest)
		return
	}

	_,err=services.RegistrationsSvc.CheckIfUserRegistered(user.Email)

	if err == nil{
		customErr:="user  already exists"
		http.Error(w, customErr, http.StatusBadRequest)
		return
	}

	err=services.RegistrationsSvc.RegisterUser(user)

	if err != nil{
		customErr:="failed to register user "
		http.Error(w, customErr, http.StatusBadRequest)
		return
	}

	err=services.LoginSvc.AddUser(user)

	if err != nil{
		customErr:="failed to add user details"
		http.Error(w, customErr, http.StatusBadRequest)
		return
	}

}

