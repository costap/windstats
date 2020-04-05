package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/logicnow/smp-poc/internal/webapp"
)

// Authorization: Bearer eyJrIjoiNXduUDJrbkE1UGllY2FpTFZnYnNJaUZzMEZCODRQSDkiLCJuIjoid2luZGdvZCIsImlkIjoxfQ==

type controller struct {
	t *template.Template
}

func main() {
	log.Println("WebApp started")

	cfg := webapp.ReadConfig()
	p := cfg.ListeningPort

	t := template.Must(template.ParseFiles(path.Join(cfg.DataRootPath, "template/index.html")))
	c := controller{t: t}

	http.HandleFunc("/", c.index)
	http.HandleFunc("/ready", c.ready)
	http.HandleFunc("/healthy", c.healthy)
	http.ListenAndServe(":"+strconv.Itoa(p), nil)
}

func (c *controller) index(w http.ResponseWriter, r *http.Request) {

	err := c.t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Printf("Error executing template %v", err)
		http.Error(w, "Oops, something went wrong...", http.StatusInternalServerError)
	}
}

func (c *controller) ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "READY")
}

func (c *controller) healthy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
