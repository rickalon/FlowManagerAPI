package router

import (
	"log"
	"net/http"

	"github.com/rickalon/FlowManagerAPI/internal/handlers"
	"github.com/rickalon/FlowManagerAPI/internal/repositories"

	"github.com/gorilla/mux"
)

type Router struct {
	Addr      string //ip:port
	Router    *mux.Router
	Subrouter *mux.Router
	DB        *repositories.PqDB
}

func NewRouter(addr string, DB *repositories.PqDB) *Router {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	return &Router{Addr: addr, Router: router, Subrouter: subrouter, DB: DB}
}

func (r *Router) CreateHandlersForSubrouter() {
	handler := handlers.NewHandler(r.Subrouter, r.DB)
	handler.CreateHandlers()
}

func (r *Router) StartListenAndServe() {
	log.Fatal(http.ListenAndServe(r.Addr, r.Router))
}
