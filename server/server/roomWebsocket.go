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

type websocketRoomHandler struct {
	roomStore    persistence.RoomStorable
	messageStore persistence.MessageStorable
	upgrader     websocket.Upgrader
	clientRooms  map[int][]*websocket.Conn
	messages     chan model.Message
}

func (w *websocketRoomHandler) handleConnections(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode("invalid username provided")
		return
	}

	rooms, err := w.roomStore.GetUserRooms(username)
	if err != nil {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(err.Error())
		return
	}

	conn, err := w.upgrader.Upgrade(rw, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for _, room := range rooms.Items {
		w.clientRooms[room.Id] = append(w.clientRooms[room.Id], conn)
	}

	fmt.Println("connected ", conn.LocalAddr())

	for {
		var msg model.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)

			for _, room := range rooms.Items {
				if len(w.clientRooms[room.Id]) == 1 {
					delete(w.clientRooms, room.Id)
				} else {
					for i, connection := range w.clientRooms[room.Id] {
						if connection == conn {
							w.clientRooms[room.Id] = append(w.clientRooms[room.Id][:i], w.clientRooms[room.Id][i+1:]...)
						}
					}
				}
			}

			return
		}

		_, err = w.messageStore.AddMessage(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		w.messages <- msg
	}
}

func (w *websocketRoomHandler) HandleMessages() {
	for {
		msg := <-w.messages

		for i, client := range w.clientRooms[msg.RoomId] {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				w.clientRooms[msg.RoomId] = append(w.clientRooms[msg.RoomId][:i], w.clientRooms[msg.RoomId][i+1:]...)
			}
		}
	}
}
