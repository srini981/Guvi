package controllers

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"guviTask/api/controllers/models"
	"guviTask/api/services"
	"io"
	"net/http"
	"time"
	"github.com/google/uuid"

)



func UserLogin(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
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

	err=services.LoginSvc.CheckCredentials(user)

	if err!=nil{
		 customErr:="user not founc"
		 http.Error(w, customErr, http.StatusNotFound)
		 return
	}
	id := uuid.New()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	err=client.Set(context.TODO(),id.String(),user.Email,time.Hour*48).Err()
	if err!=nil{
		customErr:="failed to set session"
		http.Error(w, customErr, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	obj,err:=json.Marshal(id.String())
	w.Write(obj)
}
