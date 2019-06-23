package app

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rjohnt/SampleGoApp/models"
	"github.com/rjohnt/SampleGoApp/utils"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		unauthenticatedEndpoints := []string{"/api/user/new", "/api/user/login"}
		requestPath := request.URL.Path

		// Serve call without auth if unauthenticated endpoint.
		for _, value := range unauthenticatedEndpoints {
			if value == requestPath {
				next.ServeHTTP(writer, request)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := request.Header.Get("Authorization")

		// If Token Header missing, return 403 Forbidden .
		if tokenHeader == "" {
			response = utils.Message(false, "Missing Auth Token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}

		splitTokenHeaders := strings.Split(tokenHeader, " ")

		// If Token Header improperly formatted, return 403 Forbidden .
		if len(splitTokenHeaders) != 2 {
			response = utils.Message(false, "Invalid/Malformed Auth Token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}

		tokenPart := splitTokenHeaders[1]
		token := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, token, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		// If Token malformed, return 403 Forbidden.
		if err != nil {
			response = utils.Message(false, "Malformed Authentication Token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}

		// If Token invalid, return 403 Forbidden.
		if !token.Valid {
			response = utils.Message(false, "Token is not valid.")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}

		// All checks good, proceed with request.
		fmt.Sprintf("User %", token.Username)
		ctx := context.WithValue(request.Context(), "user", token.UserId)
		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	})
}
