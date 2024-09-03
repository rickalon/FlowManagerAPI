package handlers

import (
	"github.com/rickalon/FlowManagerAPI/internal/repositories"
	"github.com/rickalon/FlowManagerAPI/internal/services"

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
	service := services.NewService(handler.Router, handler.DB)
	handler.Router.HandleFunc("/user/register", service.RegisterUser).Methods("POST")
	handler.Router.HandleFunc("/user/login", service.LoginUser).Methods("POST")
	//edit projects

	//edit tasks
}
