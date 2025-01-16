package main

import (
	"log"
	"net/http"

	"github.com/evgshul/person_g/internal/entity/cmd/config"

	"github.com/gorilla/mux"
)

func main() {

	db := config.InitDB()
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Welcome to the GoLang CRUD API!"))
		if err != nil {
			return
		}
	})

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
