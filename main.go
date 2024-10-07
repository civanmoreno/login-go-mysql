package main

import (
	"log"
	"main/auth"
	application "main/internal/users/application"
	database "main/internal/users/infrastructure/database"
	userHttp "main/internal/users/infrastructure/http"
	"main/utilities"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	/*
		pass := "123456"
		hash, err := utilities.HashPassword(pass)
		if err != nil {
			log.Fatal("Error hashing password: ", err)
		}
		println(hash)
	*/
	db, err := utilities.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	route := mux.NewRouter()

	userRepository := database.NewMySQLUserRepository(db)
	userService := application.NewUserService(userRepository)
	userHandler := userHttp.NewUserHandler(userService)

	route.HandleFunc("/user/login", userHandler.Login).Methods("POST")
	route.Handle("/user/create", auth.AuthMiddleware(http.HandlerFunc(userHandler.CreateUser))).Methods("POST")
	route.Handle("/auth/validation", auth.AuthMiddleware(http.HandlerFunc(auth.ValidateTokenHandler))).Methods("POST")

	http.ListenAndServe(":8080", route)
}
