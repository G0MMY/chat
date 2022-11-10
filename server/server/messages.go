package server

import (
	"encoding/json"
	"net/http"

	"github.com/G0MMY/chat/model"
	"github.com/G0MMY/chat/persistence"
	"github.com/gorilla/mux"
)

type messageHandler struct {
	store persistence.MessageStorable
}

func (h *messageHandler) AddMessage(w http.ResponseWriter, r *http.Request) {
	var body model.Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	defer r.Body.Close()

	msg, err := h.store.AddMessage(&body)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func (h *messageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomId, ok := vars["roomId"]
	if !ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid id provided")
		return
	}

	messages, err := h.store.GetMessages(roomId)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}
