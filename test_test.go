package main

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/yelarys4/GolangUniversity/app/handlers"
	"github.com/yelarys4/GolangUniversity/app/repositories"
	"github.com/yelarys4/GolangUniversity/app/services"
	"github.com/yelarys4/GolangUniversity/app/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// UNIT TEST
func TestGenerateToken(t *testing.T) {
	// Define test data
	userId := "123456"
	login := "testuser"
	role := "admin"
	var secretKey = []byte("1oic2oi1ensd0a9dicw121k32aspdojacs")

	tokenString, err := utils.GenerateToken(userId, login, role)
	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}

	token, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		// Return the secret key
		return secretKey, nil
	})
	if parseErr != nil {
		t.Errorf("Error parsing token: %v", parseErr)
	}

	// Validate token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		t.Errorf("Token validation failed")
	}

	// Check if userId, login, and role are correct
	if claims["userId"] != userId {
		t.Errorf("Expected userId to be %s, got %v", userId, claims["userId"])
	}
	if claims["login"] != login {
		t.Errorf("Expected login to be %s, got %v", login, claims["login"])
	}
	if claims["role"] != role {
		t.Errorf("Expected role to be %s, got %v", role, claims["role"])
	}

	// Check expiration
	exp := claims["exp"].(float64)
	expTime := time.Unix(int64(exp), 0)
	if expTime.Before(time.Now()) {
		t.Errorf("Token expired")
	}
}

// INTEGRATING TEST
func TestRegisterIntegration(t *testing.T) {

	data := map[string]interface{}{
		"login":    "test@gmail.com",
		"password": "12345",
	}
	jsonData, _ := json.Marshal(data)

	request, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(jsonData))
	response := httptest.NewRecorder()

	mongoURI := "mongodb+srv://client:5423@golang.fcwced4.mongodb.net/"
	client, _ := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	authHandler := handlers.NewAuthHandler(
		services.NewAuthService(
			repositories.NewUserRepository(client),
		),
	)

	authHandler.RegisterHandler(response, request)

	if response.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect status code. Expected: %d, Got: %d", http.StatusOK, response.Code)
	}
	expected := `{"error": "Error creating user"}`

	expectedTrimmed := strings.ReplaceAll(strings.TrimSpace(expected), " ", "")
	actualTrimmed := strings.ReplaceAll(strings.TrimSpace(response.Body.String()), " ", "")

	if actualTrimmed != expectedTrimmed {
		t.Errorf("Incorrect response body. Expected: %s, Got: %s", expected, response.Body.String())
	}
}
