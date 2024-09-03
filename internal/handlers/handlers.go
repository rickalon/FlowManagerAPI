package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/rickalon/FlowManagerAPI/internal/domain"
	"github.com/rickalon/FlowManagerAPI/internal/repositories"
	"github.com/rickalon/FlowManagerAPI/internal/services"
	"github.com/rickalon/FlowManagerAPI/pkg/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	DB     *repositories.PqDB
}

func NewHandler(router *mux.Router, DB *repositories.PqDB) *Handler {
	return &Handler{Router: router, DB: DB}
}

func (handler *Handler) CreateHandlers() {
	//login and register
	handler.Router.HandleFunc("/user/register", handler.registerUser).Methods("POST")
	handler.Router.HandleFunc("/user/login", handler.loginUser).Methods("POST")
	//edit projects

	//edit tasks
}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
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

	if err = services.ValidateUser(user); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	if err = services.CreateUser(h.DB, user); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}
	utils.WriteJSON(w, http.StatusAccepted, user)

}

func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {

}
