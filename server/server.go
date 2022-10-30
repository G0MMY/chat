package server

import (
	"net/http"

	"github.com/G0MMY/chat/persistence"
	"github.com/gorilla/mux"
)

func CreateRoutes(conn *persistence.Connection) *mux.Router {
	userHandler := userHandler{store: persistence.NewUserStore(conn)}

	router := mux.NewRouter()

	router.HandleFunc("/user", userHandler.AddUser).Methods(http.MethodPost)

	return router
}
