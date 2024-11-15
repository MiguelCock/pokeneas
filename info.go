package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
)

func info(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(pokeneas))
	randomPokenea := pokeneas[randomIndex]

	containerID := os.Getenv("HOSTNAME")
	response := struct {
		Pokenea
		ContainerID string `json:"container_id"`
	}{
		Pokenea:     randomPokenea,
		ContainerID: containerID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}