package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rickalon/FlowManagerAPI/internal/domain"
	"github.com/rickalon/FlowManagerAPI/internal/middleware"
	"github.com/rickalon/FlowManagerAPI/internal/repositories"
	"github.com/rickalon/FlowManagerAPI/pkg/utils"
)

type Service struct {
	Router *mux.Router
	DB     *repositories.PqDB
}

type TokenJWT struct {
	Token string `json:"token"`
}

func NewService(router *mux.Router, DB *repositories.PqDB) *Service {
	return &Service{Router: router, DB: DB}
}

func (service *Service) RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Registering user")
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error reading request.")
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error reading the request"})
		return
	}
	var user *domain.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error unmarshaling the body.")
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error reading the content"})
		return
	}

	if err = domain.ValidateUser(user); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}
	user.Password = hashPassword

	if err = domain.CreateUser(service.DB, user); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	if err = domain.GetIdUser(service.DB, user); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	log.Println("Token generation")
	tokenString, err := middleware.CreateTokenJWTCookie(w, user.Id)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	token := &TokenJWT{Token: tokenString}
	utils.WriteJSON(w, http.StatusAccepted, token)

}

func (service *Service) LoginUser(w http.ResponseWriter, r *http.Request) {

}
