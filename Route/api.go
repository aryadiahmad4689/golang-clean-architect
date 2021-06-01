package route

import (
	"fmt"
	infrasructure "golang-clean-architecture/Infrasructure"
	"net/http"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/add", infrasructure.AddUser)
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	return mux
}
