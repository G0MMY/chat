package server

import (
	"net/http"

	"github.com/G0MMY/chat/model"
	"github.com/G0MMY/chat/persistence"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func CreateRoutes(conn *persistence.Connection) (*mux.Router, websocketRoomHandler, websocketInvitationHandler) {
	roomStore := persistence.NewRoomStore(conn)
	messageStore := persistence.NewMessageStore(conn)
	invitationStore := persistence.NewInvitationStore(conn)

	userHandler := userHandler{store: persistence.NewUserStore(conn)}
	roomHandler := roomHandler{store: roomStore}
	invitationHandler := invitationHandler{store: invitationStore}
	messageHandler := messageHandler{store: messageStore}
	websocketRoomHandler := websocketRoomHandler{
		roomStore:    roomStore,
		messageStore: messageStore,
		upgrader:     websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
		clientRooms:  make(map[int][]*websocket.Conn),
		messages:     make(chan model.Message),
	}
	websocketInvitationHandler := websocketInvitationHandler{
		invitationStore:   invitationStore,
		upgrader:          websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
		clientInvitations: make(map[string]*websocket.Conn),
		invitations:       make(chan model.Invitation),
	}

	router := mux.NewRouter()

	roomRouter := router.PathPrefix("/rooms").Subrouter()
	invitationRouter := router.PathPrefix("/invitations").Subrouter()
	messageRouter := router.PathPrefix("/messages").Subrouter()
	websocketRouter := router.PathPrefix("/ws").Subrouter()

	roomRouter.Use(authentication)
	invitationRouter.Use(authentication)
	messageRouter.Use(authentication)
	//websocketRouter.Use(authentication)

	// user routes
	router.HandleFunc("/user", userHandler.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/validate", ValidateToken).Methods(http.MethodGet)

	// room routes
	roomRouter.HandleFunc("", roomHandler.AddRoom).Methods(http.MethodPost)
	roomRouter.HandleFunc("", roomHandler.GetUserRooms).Methods(http.MethodGet)
	roomRouter.HandleFunc("/{id}", roomHandler.GetRoomUsers).Methods(http.MethodGet)
	roomRouter.HandleFunc("/join", roomHandler.JoinRoom).Methods(http.MethodPost)

	// invitation routes
	invitationRouter.HandleFunc("", invitationHandler.AddInvitation).Methods(http.MethodPost)
	invitationRouter.HandleFunc("/{id}", invitationHandler.DeleteInvitation).Methods(http.MethodDelete)
	invitationRouter.HandleFunc("/{username}", invitationHandler.GetInvitationsOfUser).Methods(http.MethodGet)

	// message routes
	messageRouter.HandleFunc("", messageHandler.AddMessage).Methods(http.MethodPost)
	messageRouter.HandleFunc("/{roomId}", messageHandler.GetMessages).Methods(http.MethodGet)

	// websocket routes
	websocketRouter.HandleFunc("/rooms/{username}", websocketRoomHandler.handleConnections).Methods(http.MethodGet)
	websocketRouter.HandleFunc("/invitations/{username}", websocketInvitationHandler.handleConnections).Methods(http.MethodGet)

	return router, websocketRoomHandler, websocketInvitationHandler
}
