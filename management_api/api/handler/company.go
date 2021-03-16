package handler

import (
	"io"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func createCompany() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Teste")
	})
}

func MakeCompanyHandlers(router *mux.Router, n *negroni.Negroni) {
	router.Handle("/v1/company", n.With(
		negroni.Wrap(createCompany()),
	)).Methods("POST", "OPTIONS").Name("createCompany")
}
