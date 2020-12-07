package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Up .
func Up() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Deploy: Kubernetes ")
	})

	router.HandleFunc("/who", who).Methods("GET")
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":1323", handler))
}
func who(w http.ResponseWriter, r *http.Request) {
	respuesta := map[string]interface{}{
		"hostname": os.Getenv("HOSTNAME"),
		"time":     time.Now(),
	}

	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)

}
