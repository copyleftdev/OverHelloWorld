package api

import (
	"encoding/json"
	"net/http"
	"overhelloworld/internal/app"
)

type HelloRequest struct {
	Message string `json:"message"`
}

type HelloAPI struct {
	CommandHandler app.CommandHandler
}

func (api *HelloAPI) SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	var req HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	cmd := app.SayHelloCommand{Message: req.Message}
	if err := api.CommandHandler.Handle(cmd); err != nil {
		http.Error(w, "Failed to handle command", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	if _, err := w.Write([]byte("Hello event accepted!")); err != nil {
		// log error or handle as needed
		_ = err
	}
}
