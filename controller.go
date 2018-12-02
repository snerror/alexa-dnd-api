package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	buildResponse(w, r, "player created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func PlayerAttackEnemy(w http.ResponseWriter, r *http.Request) {
	var enemy Enemy
	var ability Ability

	vars := mux.Vars(r)
	a := vars["ability"]
	enemyId, err := strconv.Atoi(vars["enemyId"])
	if err != nil {
		serverError(w, r, err)
		return
	}

	for _, e := range enemies {
		if e.ID == enemyId {
			enemy = e
		}
	}

	if enemy.Name == "" {
		buildResponse(w, r, "unknown enemy targeted")
	}

	for _, ab := range player.Abilities {
		if strings.ToLower(ab.Name) == strings.ToLower(a) {
			ability = ab
		}
	}

	if ability.Name == "" {
		buildResponse(w, r, "unknown ability used")
	}

	response := player.AttackEnemy(ability, enemy)
	buildResponse(w, r, response)
}

func buildResponse(w http.ResponseWriter, r *http.Request, value string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"value": value,
	}); err != nil {
		serverError(w, r, err)
		return
	}
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
