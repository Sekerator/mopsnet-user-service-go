package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"user/internal/db/entities"
	"user/internal/services"
)

type UserHand struct {
	userServ services.UserServices
}

func NewUserHandler(userServ services.UserServices) UserHandler {
	return &UserHand{userServ}
}

func (uh *UserHand) Login(w http.ResponseWriter, r *http.Request) {
	userData := entities.UserJson{}

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		userData.Username = r.FormValue("username")
		userData.Password = r.FormValue("password")
		userData.Email = r.FormValue("email")
	} else {
		http.Error(w, "Unsupported content type", http.StatusUnsupportedMediaType)
		return
	}

	if userData.Username == "" || userData.Password == "" {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	user, err := uh.userServ.Login(userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	userJson := entities.UserReturnJson{
		Username:  user.Username,
		Email:     user.Email,
		AuthToken: user.AuthToken,
		CreatedAt: user.CreatedAt,
	}

	err = json.NewEncoder(w).Encode(userJson)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}
}
