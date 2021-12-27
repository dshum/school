package auth

import (
	"github.com/dshum/school/internal/utils"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type Service interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type authService struct {
	storage Storage
	redis   *redis.Client
}

func NewService(storage Storage, redis *redis.Client) Service {
	return &authService{
		storage: storage,
		redis:   redis,
	}
}

func (s *authService) Login(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	user, attemptErr := s.storage.Attempt(login, password)
	if attemptErr != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": attemptErr.Error()})
		return
	}

	token, tokenErr := utils.CreateToken(user.Id)
	if tokenErr != nil {
		utils.JSONResponse(w, http.StatusUnprocessableEntity, map[string]string{"error": tokenErr.Error()})
		return
	}

	// authErr := utils.CreateAuth(user.Id, token)
	// if authErr != nil {
	// 	utils.JSONResponse(w, http.StatusUnprocessableEntity, map[string]string{"error": authErr.Error()})
	// 	return
	// }

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{"id": user.Id, "token": token})
}

func (s *authService) Logout(w http.ResponseWriter, _ *http.Request) {
	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Logged out"})
}
