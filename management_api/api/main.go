package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/marcelocpinheiro/communication-engine/infrastructure/repository"
	"github.com/marcelocpinheiro/communication-engine/management_api/api/handler"
	"github.com/marcelocpinheiro/communication-engine/management_api/config"
	"github.com/marcelocpinheiro/communication-engine/usecase/company"
)

func main() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	fmt.Println(dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	companyRepository := repository.NewCompanyMySQL(db)
	companyService := company.NewService(companyRepository)

	r := mux.NewRouter()
	n := negroni.New()

	handler.MakeCompanyHandlers(r, n, companyService)

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

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
