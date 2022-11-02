package server

import (
	"encoding/json"
	"net/http"

	"github.com/G0MMY/chat/model"
	"github.com/G0MMY/chat/persistence"
	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
)

type invitationHandler struct {
	store persistence.InvitationStorable
}

func (h *invitationHandler) AddInvitation(w http.ResponseWriter, r *http.Request) {
	var body model.Invitation

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	defer r.Body.Close()

	invitation, err := h.store.AddInvitation(&body)
	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invitation)
}

func (h *invitationHandler) GetInvitationsOfUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid username provided")
		return
	}

	room, err := h.store.GetInvitationsOfUser(username)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(room)
}
