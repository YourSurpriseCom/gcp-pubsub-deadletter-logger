package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger/cloud"
	"github.com/gorilla/mux"
	log "github.com/jlentink/yaglogger"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Read the body to a byte array for better processing
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil || len(bodyBytes) == 0 {
			log.Warn("Unable to read body and convert it to byte array: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Log full message
		cloud.LogFullMessage(bodyBytes)

		//Log Error with normalized data
		cloud.LogError(bodyBytes)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok!")
	}).Methods(http.MethodPost)

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "pong")
	}).Methods(http.MethodGet)

	return router
}
