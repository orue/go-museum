package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"orue.io/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	if id != nil {
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) { // Get one exhibition
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		}

	} else { //Get All exhibitions
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
