package http

import (
	"encoding/json"
	"main/internal/users/application"
	"main/internal/users/domain"
	"main/utilities"
	"net/http"
)

type UserHandler struct {
	Service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "All fields (username, email, password, role) are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utilities.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user.Password = string(hashedPassword)

	err = h.Service.CreateUser(user)
	if err != nil {
		http.Error(w, "error al crear el usuario", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var userJSON struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	if creds.Email == "" || creds.Password == "" {
		http.Error(w, "Missing email or password", http.StatusBadRequest)
		return
	}
	user, token, err := h.Service.Login(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userJSON.ID = user.ID
	userJSON.Username = user.Username
	userJSON.Email = user.Email
	userJSON.Token = token

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userJSON)
}

func (h *UserHandler) RessetPassword(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}

	err = h.Service.RessetPassword(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
