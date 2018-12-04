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

func SetPlayerClassAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	className := vars["className"]

	err := player.CreateFromTemplate(className)
	if err != nil {
		serverError(w, r, err)
		return
	}

	textResponse := "You've chosen " + player.Name + " the " + player.Class + ". To know more about your class ask me 'what are my stats?'."

	buildResponse(w, r, textResponse)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func GetPlayerDetailsAction(w http.ResponseWriter, r *http.Request) {
	if player.Name == "" {
		buildResponse(w, r, "You must choose a class first")
		return
	}

	textResponse := "You play a " + player.Class + " named " + player.Name + ". Your current HP is " + strconv.Itoa(player.CurrentHp) + " out of " + strconv.Itoa(player.MaxHp) + " and your armor class is " + strconv.Itoa(player.ArmorClass) + ". You also have " + strconv.Itoa(len(player.Abilities)) + " abilities."

	for i := 0; i < len(player.Abilities); i++ {
		textResponse += player.Abilities[i].Name + " has attack of " + strconv.Itoa(player.Abilities[i].Attack) + " and damage of " + strconv.Itoa(player.Abilities[i].Damage) + "."

		if player.Abilities[i].CD != 0 {
			textResponse += "It also has a cooldown of " + strconv.Itoa(player.Abilities[i].CD) + "."
		}
	}

	buildResponse(w, r, textResponse)
}

func PlayerAttackEnemyAction(w http.ResponseWriter, r *http.Request) {
	var enemy *Enemy
	var ability *Ability

	vars := mux.Vars(r)
	a := vars["ability"]
	enemyId, err := strconv.Atoi(vars["enemyId"])
	if err != nil {
		serverError(w, r, err)
		return
	}

	for i := 0; i < len(enemies); i++ {
		if enemies[i].ID == enemyId {
			enemy = &enemies[i]
		}
	}

	if enemy == nil {
		buildResponse(w, r, "enemy not found")
		return
	}

	for i := 0; i < len(player.Abilities); i++ {
		if strings.ToLower(player.Abilities[i].Name) == strings.ToLower(a) {
			ability = &player.Abilities[i]
		}

		if player.Abilities[i].CurrentCD > 0 {
			player.Abilities[i].CurrentCD--
		}
	}

	if ability == nil {
		buildResponse(w, r, "unknown ability used")
		return
	}

	buildResponse(w, r, player.AttackEnemy(ability, enemy))
}

func MovePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	direction := vars["direction"]

	buildResponse(w, r, player.Move(&dungeon, direction))
}

func EnemyAttackPlayerAction(w http.ResponseWriter, r *http.Request) {
	var enemy Enemy

	vars := mux.Vars(r)
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
		buildResponse(w, r, "unknown enemy attacked")
	}

	buildResponse(w, r, enemy.AttackPlayer())
}

func buildResponse(w http.ResponseWriter, r *http.Request, value string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"alexaResponse": value,
		"player":        player,
		"enemies":       enemies,
	}); err != nil {
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
