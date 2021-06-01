package main

import (
	route "golang-clean-architecture/Route"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := route.Route()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
