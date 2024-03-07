package routes

import (
	"github.com/gorilla/mux"
	"github.com/yelarys4/GolangUniversity/app/handlers"
	"github.com/yelarys4/GolangUniversity/app/repositories"
	"github.com/yelarys4/GolangUniversity/app/services"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func SetRoutes(router *mux.Router, client *mongo.Client) {
	authHandler := handlers.NewAuthHandler(
		services.NewAuthService(
			repositories.NewUserRepository(client),
		),
	)

	router.HandleFunc("/auth/delete", authHandler.DeleteHandler).Methods("POST")
	router.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", authHandler.LogoutHandler).Methods("GET")
	router.HandleFunc("/auth/register", authHandler.RegisterHandler).Methods("POST")
	router.HandleFunc("/auth/verify", authHandler.VerifyHandler).Methods("GET")
	router.HandleFunc("/books/", handlers.GetPaginatedItems).Methods("GET")
	router.HandleFunc("/add_book", handlers.AddBookHandler).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
}
