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

	text := "You've chosen " + player.Name + " the " + player.Class + ". To know more about your class ask me what are my stats?. You can now begin your journey. "

	buildResponse(w, r, text)
}

func GetPlayerDetailsAction(w http.ResponseWriter, r *http.Request) {
	if isPlayerCreated() == false {
		buildResponse(w, r, "You must choose a class first")
		return
	}

	text := "You play a " + player.Class + " named " + player.Name + ". Your current HP is " + strconv.Itoa(player.CurrentHp) + " out of " + strconv.Itoa(player.MaxHp) + " and your armor class is " + strconv.Itoa(player.ArmorClass) + ". You also have " + strconv.Itoa(len(player.Abilities)) + " abilities."

	for i := 0; i < len(player.Abilities); i++ {
		text += " " + player.Abilities[i].Name + " has attack of " + strconv.Itoa(player.Abilities[i].Attack) + " and damage of " + strconv.Itoa(player.Abilities[i].Damage) + ". "

		if player.Abilities[i].CD != 0 {
			text += " It also has a cooldown of " + strconv.Itoa(player.Abilities[i].CD) + ". "
		}
	}

	buildResponse(w, r, text)
}

func PlayerAttackAction(w http.ResponseWriter, r *http.Request) {
	if isPlayerCreated() == false {
		buildResponse(w, r, "You must choose a class first")
		return
	}

	var e *Enemy
	var a *Ability

	vars := mux.Vars(r)
	v := vars["ability"]

	for i := 0; i < len(enemies); i++ {
		if enemies[i].Position.X == player.Position.X && enemies[i].Position.Y == player.Position.Y {
			e = &enemies[i]
		}
	}

	if e == nil {
		buildResponse(w, r, "No enemies around you.")
		return
	}

	for i := 0; i < len(player.Abilities); i++ {
		if strings.ToLower(player.Abilities[i].Name) == strings.ToLower(v) {
			a = &player.Abilities[i]
		}

		if player.Abilities[i].CurrentCD > 0 {
			player.Abilities[i].CurrentCD--
		}
	}

	if a == nil {
		buildResponse(w, r, "You don't have that ability.")
		return
	}

	playerAttackText := player.AttackEnemy(a, e)
	enemyAttackText := ""
	if e.CurrentHp > 0 {
		enemyAttackText = e.AttackPlayer()
	} else {
		for i := 0; i < len(enemies); i++ {
			if &enemies[i] == e {
				if i == len(enemies)-1 {
					enemies = enemies[:len(enemies)-1]
				} else {
					enemies = append(enemies[:i], enemies[i+1:]...)
				}
			}
		}
	}

	buildResponse(w, r, playerAttackText+" "+enemyAttackText)
}

func MovePlayer(w http.ResponseWriter, r *http.Request) {
	if isPlayerCreated() == false {
		buildResponse(w, r, "You must choose a class first")
		return
	}
	var e *Enemy

	vars := mux.Vars(r)
	d := vars["direction"]

	for i := 0; i < len(enemies); i++ {
		if enemies[i].Position.X == player.Position.X && enemies[i].Position.Y == player.Position.Y {
			e = &enemies[i]
		}
	}

	if e != nil {
		buildResponse(w, r, " Unable to move, there is a "+e.Name+" blocking your way. ")
		return
	}

	if _, err := player.Move(&dungeon, d); err != nil {
		buildResponse(w, r, fmt.Sprintf("%s", err))
		return
	}
	enemyDetectedText := ""

	for i := 0; i < len(enemies); i++ {
		if enemies[i].Position.X == player.Position.X && enemies[i].Position.Y == player.Position.Y {
			e = &enemies[i]
		}
	}

	if e != nil {
		enemyDetectedText = " A " + e.Name + " emerges from the dark."
	}

	availablePathsText := " Available paths are "
	if dungeon.cells[player.Position.X][player.Position.Y]&up != 0 {
		availablePathsText += " up,"
	}

	if dungeon.cells[player.Position.X][player.Position.Y]&down != 0 {
		availablePathsText += " down,"
	}

	if dungeon.cells[player.Position.X][player.Position.Y]&right != 0 {
		availablePathsText += " right,"
	}

	if dungeon.cells[player.Position.X][player.Position.Y]&left != 0 {
		availablePathsText += " left,"
	}
	availablePathsText += "."

	buildResponse(w, r, "You moved. "+availablePathsText+enemyDetectedText)
}

func PlayerAbilitiesAction(w http.ResponseWriter, r *http.Request) {
	if isPlayerCreated() == false {
		buildResponse(w, r, "You must choose a class first")
		return
	}
	text := ""

	for i := 0; i < len(player.Abilities); i++ {
		text += " " + player.Abilities[i].Name + " has attack of " + strconv.Itoa(player.Abilities[i].Attack) + " and damage of " + strconv.Itoa(player.Abilities[i].Damage) + ". "

		if player.Abilities[i].CD != 0 {
			text += " It also has a cooldown of " + strconv.Itoa(player.Abilities[i].CD)
		}

		if player.Abilities[i].CurrentCD != 0 {
			text += " and can be used again in " + strconv.Itoa(player.Abilities[i].CurrentCD) + " turns"
		}
	}

	text += "."

	buildResponse(w, r, text)
}

func ResetAction(w http.ResponseWriter, r *http.Request) {
	enemies = []Enemy{}
	player = Player{}
	dungeon = CreateDungeon(10, 30)
	dungeon.generate()

	buildResponse(w, r, "Dungeon reset successfully.")
}

func isPlayerCreated() bool {
	if player.Name != "" {
		return true
	}

	return false
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
