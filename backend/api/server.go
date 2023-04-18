package api

import (
	"fmt"
	_ "github.com/lib/pq" // <------------ here
	"guviTask/api/controllers"
	"guviTask/api/middleware"
	"log"
	"net/http"
)

func Run(){
	//http.Handle("/", corsMiddleware(myHandler))

	http.HandleFunc("/signup",middleware.CorsMiddleware( controllers.RegisterUser))
	http.HandleFunc("/login", middleware.CorsMiddleware(controllers.UserLogin))
	http.HandleFunc("/profile", middleware.AuthMiddleware(controllers.UserProfile))
	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

