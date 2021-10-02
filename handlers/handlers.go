package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		log.Print(r.Body)
		fmt.Fprintf(w, "Welcome to GO REST API")
	})

	//Parameter - APM Formula
	myRouter.HandleFunc("/apm-formula", GetAllApmFormula).Methods("GET")
	// myRouter.HandleFunc("/apm-formula/add", StoreApmFormula).Methods("POST")
	myRouter.HandleFunc("/apm-formula/{metric_id}", GetApmFormulaById).Methods("GET")
	myRouter.HandleFunc("/apm-formula", UpdateApmFormula).Methods("PUT", "POST")
	myRouter.HandleFunc("/apm-formula", DestroyApmFormula).Methods("DELETE")

	handler := cors.Default().Handler(myRouter)

    log.Fatal(http.ListenAndServe(":11000", handler))
}