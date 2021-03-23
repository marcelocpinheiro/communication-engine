package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/marcelocpinheiro/communication-engine/management_api/api/presenter"
	"github.com/marcelocpinheiro/communication-engine/usecase/company"
)

func createCompany(service *company.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding company"
		var input struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := service.CreateCompany(input.Name, input.Email)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJson := &presenter.Company{
			ID:    id,
			Email: input.Email,
			Name:  input.Name,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func MakeCompanyHandlers(router *mux.Router, n *negroni.Negroni, service *company.Service) {
	router.Handle("/v1/company", n.With(
		negroni.Wrap(createCompany(service)),
	)).Methods("POST", "OPTIONS").Name("createCompany")
}
