package server

import (
	"net/http"

	"github.com/G0MMY/chat/persistence"
	"github.com/gorilla/mux"
)

func CreateRoutes(conn *persistence.Connection) *mux.Router {
	userHandler := userHandler{store: persistence.NewUserStore(conn)}
	roomHandler := roomHandler{store: persistence.NewRoomStore(conn)}

	router := mux.NewRouter()

	// user routes
	router.HandleFunc("/user", userHandler.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)

	// room routes
	router.HandleFunc("/room", roomHandler.AddRoom).Methods(http.MethodPost)
	router.HandleFunc("/room/{name}", roomHandler.GetRoom).Methods(http.MethodGet)
	router.HandleFunc("/room/join", roomHandler.JoinRoom).Methods(http.MethodPost)

	return router
}
