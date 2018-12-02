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
		"Get",
		"/player",
		GetPlayerClassAction,
	},
	Route{
		"SetPlayerClass",
		"POST",
		"/player/class/{className}",
		SetPlayerClassAction,
	},
	Route{
		"PlayerAttackEnemy",
		"GET",
		"/player/attack/{ability}/{enemyId}",
		PlayerAttackEnemy,
	},
}
