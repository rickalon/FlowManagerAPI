package handlers

import (
	"github.com/rickalon/FlowManagerAPI/internal/middleware"
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
	//projects
	handler.Router.HandleFunc("/proyect", middleware.ValidateJWT(service.CreateProyect, service)).Methods("POST")
	handler.Router.HandleFunc("/proyect/{id}", middleware.ValidateJWT(service.GetProyect, service)).Methods("GET")
	handler.Router.HandleFunc("/proyect/{id}", middleware.ValidateJWT(service.DeleteProyect, service)).Methods("DELETE")
	//tasks
}
