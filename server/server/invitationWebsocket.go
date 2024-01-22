package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/G0MMY/chat/model"
	"github.com/G0MMY/chat/persistence"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type websocketInvitationHandler struct {
	invitationStore   persistence.InvitationStorable
	upgrader          websocket.Upgrader
	clientInvitations map[string]*websocket.Conn
	invitations       chan model.Invitation
}

func (w *websocketInvitationHandler) handleConnections(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode("invalid username provided")
		return
	}

	conn, err := w.upgrader.Upgrade(rw, r, nil)
	if err != nil {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(err.Error())
		return
	}
	defer conn.Close()

	w.clientInvitations[username] = conn

	for {
		var invitation model.Invitation
		err := conn.ReadJSON(&invitation)
		if err != nil {
			fmt.Println(err)

			delete(w.clientInvitations, username)

			return
		}

		_, err = w.invitationStore.AddInvitation(&invitation)
		if err != nil {
			fmt.Println(err)
			return
		}

		w.invitations <- invitation
	}
}

func (w *websocketInvitationHandler) HandleInvitations() {
	for {
		invitation := <-w.invitations

		if conn, ok := w.clientInvitations[invitation.Receiver]; ok {
			err := conn.WriteJSON(invitation)
			if err != nil {
				fmt.Println(err)
				conn.Close()
				delete(w.clientInvitations, invitation.Receiver)
			}
		}
	}
}
