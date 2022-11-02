package server

import (
	"net/http"

	"github.com/G0MMY/chat/persistence"
	"github.com/gorilla/mux"
)

func CreateRoutes(conn *persistence.Connection) *mux.Router {
	userHandler := userHandler{store: persistence.NewUserStore(conn)}
	roomHandler := roomHandler{store: persistence.NewRoomStore(conn)}
	invitationHandler := invitationHandler{store: persistence.NewInvitationStore(conn)}

	router := mux.NewRouter()

	// user routes
	router.HandleFunc("/user", userHandler.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)

	// room routes
	router.HandleFunc("/room", roomHandler.AddRoom).Methods(http.MethodPost)
	router.HandleFunc("/room/{id}", roomHandler.GetRoomUsers).Methods(http.MethodGet)
	router.HandleFunc("/room/join", roomHandler.JoinRoom).Methods(http.MethodPost)

	// invitation routes
	router.HandleFunc("/invitation", invitationHandler.AddInvitation).Methods(http.MethodPost)
	router.HandleFunc("/invitation/{username}", invitationHandler.GetInvitationsOfUser).Methods(http.MethodGet)

	return router
}
