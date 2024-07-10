package main

import (
	"log"
	"net/http"

	"github.com/akshay0074700747/my-sandbox/controllers"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/executeCode", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("the requested service is not found"))
			return
		}
		controllers.ExecuteCode(w, r)
	})

	log.Println("sandbox has started at port 8080")
	http.ListenAndServe(":8080", mux)
}
