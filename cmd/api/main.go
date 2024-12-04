package main

import (
	"fmt"
	"net/http"

	"cryptopals-challenge-go/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus" // alised as log
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Strarting go")
	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
