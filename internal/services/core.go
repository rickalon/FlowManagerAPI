package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rickalon/FlowManagerAPI/internal/domain"
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

type JsonInfo struct {
	Info string `json:"info"`
}

func NewService(router *mux.Router, DB *repositories.PqDB) *Service {
	return &Service{Router: router, DB: DB}
}

func (service *Service) RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Registering user")
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
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
	tokenString, err := utils.CreateTokenJWTCookie(w, user.Id)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	token := &TokenJWT{Token: tokenString}
	utils.WriteJSON(w, http.StatusAccepted, token)

}

func (service *Service) LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Login user")
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error content body of request")
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error content body of request"})
		return
	}

	var user *domain.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error unmarshaling the data")
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error unmarshaling the data"})
		return
	}

	if err = domain.ValidateUserLogin(user); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}

	strPass := user.Password
	if err = domain.GetLoginUser(service.DB, user); err != nil {
		log.Println("User doesn't exists")
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "User doesn't exists"})
		return
	}

	if err = utils.HashLogin(user.Password, strPass); err != nil {
		log.Println("Wrong Password")
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Wrong Password"})
		return
	}

	log.Println("Usuario logeado es: ", user)

	strToken, err := utils.CreateTokenJWTCookie(w, user.Id)
	if err != nil {
		log.Println("Error generating the token")
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error generating the token"})
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, &TokenJWT{Token: strToken})
}

func (service *Service) CreateProyect(w http.ResponseWriter, r *http.Request, idUser int) {
	defer r.Body.Close()
	log.Println("Creaitng proyect")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request.")
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error reading the request"})
		return
	}
	var proyect *domain.Proyect
	err = json.Unmarshal(body, &proyect)
	if err != nil {
		log.Println("Error unmarshaling the body.")
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error reading the content"})
		return
	}

	if err = domain.ValidateProyect(proyect); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	if err = domain.CreateProyect(service.DB, proyect); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	if err = domain.GetProyectByName(service.DB, proyect); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, proyect)

}

func (service *Service) GetProyect(w http.ResponseWriter, r *http.Request, idUser int) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	id := vars["id"]
	idUser, err := strconv.Atoi(id)
	log.Println("Getting proyect id: ", idUser)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	proyect := &domain.Proyect{Proyect_id: idUser}
	if err = domain.GetProyectById(service.DB, proyect); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}
	//get all the proyects
	utils.WriteJSON(w, http.StatusAccepted, proyect)

}

func (service *Service) DeleteProyect(w http.ResponseWriter, r *http.Request, idUser int) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	id := vars["id"]
	idUser, err := strconv.Atoi(id)
	log.Println("Deletting proyect id: ", idUser)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	proyect := &domain.Proyect{Proyect_id: idUser}
	//First delete all the tasks for this proyect
	if err = domain.RemoveProyect(service.DB, proyect); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}
	//get all the proyects
	utils.WriteJSON(w, http.StatusAccepted, &JsonInfo{Info: "task and proyect deleted"})

}
