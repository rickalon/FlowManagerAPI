package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rickalon/FlowManagerAPI/internal/domain"
	"github.com/rickalon/FlowManagerAPI/internal/services"

	"github.com/gorilla/mux"
)

func TestRegisterUser(t *testing.T) {

	service := &services.Service{}
	t.Run("should return an error if name is empty", func(t *testing.T) {
		payload := &domain.User{
			Name:     "",
			Password: "",
			Email:    "",
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/user/register", service.RegisterUser)

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Error("Invalid status code, it should fail")
		}

	})
	t.Run("should return an error if password is empty", func(t *testing.T) {
		payload := &domain.User{
			Name:     "Ricardo",
			Password: "",
			Email:    "",
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/user/register", service.RegisterUser)

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Error("Invalid status code, it should fail")
		}

	})
	t.Run("should return an error if email is empty", func(t *testing.T) {
		payload := &domain.User{
			Name:     "Ricardo",
			Password: "1234",
			Email:    "",
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/user/register", service.RegisterUser)

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Error("Invalid status code, it should fail")
		}

	})
}
