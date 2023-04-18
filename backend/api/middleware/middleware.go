package middleware

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func CorsMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,session-id",)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		handler(w,r)
	})
}

func AuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {

	return CorsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		sessionkey:=r.Header.Get("session-id")
		// Check if the user is authenticated
		sessionkey=sessionkey[0:len(sessionkey)-1]
		client := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		})

		val:=	client.Get(context.TODO(),sessionkey)
		if val.Val()==""{
			customErr:="session not found"
			http.Error(w, customErr, http.StatusNotFound)

			return
		}

		r.Header.Set("usermail",val.Val())

		// Call the handler function if the user is authenticated
		handler(w, r)
	})
}
	//
	//func AuthenticationMiddleWare(next http.Handler) http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		var user models.Userprofile
	//		arr,err:=io.ReadAll(r.Body)
	//
	//		if err != nil{
	//			customErr:="failed to parse request body"
	//			http.Error(w, customErr, http.StatusBadRequest)
	//			return
	//		}
	//
	//		err=json.Unmarshal(arr,&user)
	//		if err != nil{
	//			customErr:="failed to parse request body"
	//			http.Error(w, customErr, http.StatusBadRequest)
	//			return
	//		}
	//
	//		client := redis.NewClient(&redis.Options{
	//			Addr: "localhost:6379",
	//			Password: "",
	//			DB: 0,
	//		})
	//
	//		val:=	client.Get(context.TODO(),user.PhoneNumber)
	//		if val==nil{
	//			customErr:="session not found"
	//			http.Error(w, customErr, http.StatusNotFound)
	//
	//			return
	//		}
	//
	//
	////		if r.Method == "OPTIONS" {
	////	w.WriteHeader(http.StatusNoContent)
	////	return
	////}
	//
	//
	//
	//	next.ServeHTTP(w, r)
	//})
	//}
