package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func IndexAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Alexa DND!")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetPlayerClassAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(player); err != nil {
		serverError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SetPlayerClassAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	className := vars["className"]

	err := player.CreateFromTemplate(className)
	if err != nil {
		serverError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func serverError(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	if _, err = fmt.Fprintln(w, err.Error()); err != nil {
		resWriteError(w, r, err)
	}
}
func resWriteError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Response write error for %s: %s", r.URL, err)
}
