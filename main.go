package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", SquareRootRoute)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	fmt.Print(srv.ListenAndServe())
}

/*
SquareRootRoute - Calculates the square root of a number in a loop, returning the result
*/
func SquareRootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	sqrt := SquareRoot()
	message := "Code.Education Rocks! \n-------------------------\nThe number is: %v"

	fmt.Fprintf(w, message, sqrt)
}
