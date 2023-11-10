package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Avalanche(w http.ResponseWriter, r *http.Request) {}
func Tsunami(w http.ResponseWriter, r *http.Request) {}
func Metrics(w http.ResponseWriter, r *http.Request) {}
func Tests(w http.ResponseWriter, r *http.Request) {}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/avalanche", Avalanche).Methods(http.MethodPost)
    r.HandleFunc("/tsunami/{num}", Tsunami).Methods(http.MethodPost)
    r.HandleFunc("/metrics/{testId}", Metrics).Methods(http.MethodGet)
    r.HandleFunc("/tests", Tests).Methods(http.MethodGet)

    http.ListenAndServe(":8000", r)
}
