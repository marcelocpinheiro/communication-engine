package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/marcelocpinheiro/communication-engine/management_api/api/handler"
	"github.com/marcelocpinheiro/communication-engine/management_api/config"
)

func main() {

	r := mux.NewRouter()
	n := negroni.New()

	handler.MakeCompanyHandlers(r, n)

	http.Handle("/", r)

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Running")
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	log.Printf("Listening on port %d \n", config.API_PORT)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
