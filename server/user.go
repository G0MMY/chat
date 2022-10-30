package server

import (
	"encoding/json"
	"net/http"

	"github.com/G0MMY/chat/model"
	"github.com/G0MMY/chat/persistence"
)

type userHandler struct {
	store persistence.UserStorable
}

func (h *userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var body model.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	defer r.Body.Close()

	user, err := h.store.AddUser(&body)
	if err != nil {
		// if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {

		// }
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
