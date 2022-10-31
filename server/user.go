package server

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/G0MMY/chat/model"
	"github.com/G0MMY/chat/persistence"
	"github.com/G0MMY/chat/util"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
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

	hashedPwd := sha256.Sum256([]byte(body.Password))
	body.Password = fmt.Sprintf("%x", hashedPwd[:])

	user, err := h.store.AddUser(&body)
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
	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var body model.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	defer r.Body.Close()

	hashedPwd := sha256.Sum256([]byte(body.Password))
	body.Password = fmt.Sprintf("%x", hashedPwd[:])

	token, err := h.store.Login(&body)
	if err != nil {
		if err == pgx.ErrNoRows {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
	}

	fmt.Println(util.ValidateToken(token))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(token)
}
