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
	"github.com/rs/cors"
)

func main() {
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
	route.HandleFunc("/user/resset_password", userHandler.RessetPassword).Methods("POST")
	route.HandleFunc("/user/create", userHandler.CreateUser).Methods("POST")

	route.Handle("/auth/validation", auth.AuthMiddleware(http.HandlerFunc(auth.ValidateTokenHandler))).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	routes := c.Handler(route)
	http.ListenAndServe(":8080", routes)

}
