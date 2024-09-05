package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rickalon/FlowManagerAPI/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func HashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hashPass), err
}

func HashLogin(hash, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}

func CreateTokenJWTCookie(w http.ResponseWriter, id int) (string, error) {

	expires := time.Now().Add(time.Minute * 50)
	claims := jwt.MapClaims{
		"exp":     expires.Unix(),
		"user_id": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.ENV.GetJWTKey()))
	http.SetCookie(w, &http.Cookie{
		Name:    "Authorization",
		Value:   tokenString,
		Expires: expires,
	})

	return tokenString, err
}
