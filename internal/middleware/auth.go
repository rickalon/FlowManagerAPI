package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/rickalon/FlowManagerAPI/internal/config"
	"github.com/rickalon/FlowManagerAPI/internal/domain"
	"github.com/rickalon/FlowManagerAPI/internal/services"
	"github.com/rickalon/FlowManagerAPI/pkg/utils"
)

func ValidateJWT(handler http.HandlerFunc, service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		strSecret := []byte(config.ENV.GetJWTKey())
		strToken := getToken(r)
		if strToken == "" {
			log.Println("Token not in the request")
			utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "token not in the request"})
			return
		}
		token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metodo de firma no es HMAC, %v", token.Header["alg"])
			}
			return strSecret, nil
		})
		if err != nil {
			log.Println(err)
			utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if id, ok := claims["user_id"]; ok {
				err = domain.GetIdUserById(service.DB, int(id.(float64)))
				if err != nil {
					log.Println(err)
					utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
					return
				} else {
					log.Println("Token confirmado de usuario: ", int(id.(float64)))
				}
			} else {
				log.Println("Id of the token is not valid")
				utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "Id of the token is not valid"})
				return
			}
		} else { //maybe
			log.Println("Token is not valid")
			utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "Token is not valid"})
			return
		}

		handler(w, r)
	}
}

func getToken(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenURL := r.URL.Query().Get("token")
	tokenCookie, err := r.Cookie("authorization")
	if tokenAuth != "" {
		return tokenAuth
	}
	if tokenURL != "" {
		return tokenURL
	}
	if err == nil {
		return tokenCookie.Value
	}
	return ""
}
