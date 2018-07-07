package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/dietDay", getDietDaysHandler).Methods("GET")
	r.HandleFunc("/dietDay", createDietDayHandler).Methods("POST")

	return r
}

func main() {
	r := newRouter()

	http.ListenAndServe(":8080", r)
}
