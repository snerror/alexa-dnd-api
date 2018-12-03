package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

		logAppState()
	})
}

func logAppState() {
	log.Printf("PLAYER: %+v\n", player)
	log.Printf("ENEMIES: %+v\n", enemies)
	dungeon.DrawDungeon()
}
