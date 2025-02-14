package controllers

import (
	"encoding/json"
	"net/http"
)

func GetApiHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": "API is connected and running",
	})
}
