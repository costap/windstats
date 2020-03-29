package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/costap/windstats/internal/app"
	"github.com/costap/windstats/internal/service"
	"github.com/julienschmidt/httprouter"
)

var s *service.WindStatsService

func main() {
	log.Printf("windstatsd started")

	c := app.ReadConfig()

	log.Printf("staring with config %v", c)

	mc := app.NewMetClient(c.APIAddr)
	dbc := app.NewDBClient(c.DBAdrr, c.DBName, c.DBUser, c.DBPass)
	defer dbc.Close()

	s = service.NewWindStatsService(mc, dbc, c)

	go s.Run()

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/healthy", healthy)
	router.GET("/stop", stop)
	router.GET("/start", start)

	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Windstad\n")
}

func healthy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if s.Healthy() {
		fmt.Fprint(w, "OK\n")
		return
	}
	http.Error(w, "Oops", http.StatusInternalServerError)
}

func stop(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s.Stop()
	fmt.Fprint(w, "OK\n")
}

func start(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	go s.Run()
	fmt.Fprint(w, "OK\n")
}
