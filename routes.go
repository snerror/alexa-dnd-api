package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		IndexAction,
	},
	Route{
		"GetPlayerClass",
		"GET",
		"/player",
		GetPlayerDetailsAction,
	},
	Route{
		"SetPlayerClass",
		"GET",
		"/player/class/{className}",
		SetPlayerClassAction,
	},
	Route{
		"MovePlayer",
		"GET",
		"/player/move/{direction}",
		MovePlayer,
	},
	Route{
		"PlayerAttack",
		"GET",
		"/player/attack/{ability}",
		PlayerAttackAction,
	},
	Route{
		"ResetDungeon",
		"GET",
		"/reset",
		ResetAction,
	},
}
