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
	messageHandler := messageHandler{store: persistence.NewMessageStore(conn)}

	router := mux.NewRouter()

	roomRouter := router.PathPrefix("/room").Subrouter()
	invitationRouter := router.PathPrefix("/invitation").Subrouter()
	messageRouter := router.PathPrefix("/message").Subrouter()

	roomRouter.Use(Authentication)
	invitationRouter.Use(Authentication)
	messageRouter.Use(Authentication)

	// user routes
	router.HandleFunc("/user", userHandler.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)

	// room routes
	roomRouter.HandleFunc("", roomHandler.AddRoom).Methods(http.MethodPost)
	roomRouter.HandleFunc("/{id}", roomHandler.GetRoomUsers).Methods(http.MethodGet)
	roomRouter.HandleFunc("/join", roomHandler.JoinRoom).Methods(http.MethodPost)

	// invitation routes
	invitationRouter.HandleFunc("", invitationHandler.AddInvitation).Methods(http.MethodPost)
	invitationRouter.HandleFunc("/{username}", invitationHandler.GetInvitationsOfUser).Methods(http.MethodGet)

	// message routes
	messageRouter.HandleFunc("", messageHandler.AddMessage).Methods(http.MethodPost)
	messageRouter.HandleFunc("/{roomId}", messageHandler.GetMessages).Methods(http.MethodGet)

	return router
}
