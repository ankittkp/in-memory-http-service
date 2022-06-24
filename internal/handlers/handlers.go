package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Handler struct {
	Database map[string]interface{}
}

func NewHandler() *Handler {
	return &Handler{
		Database: map[string]interface{}{
			"abc-1": 100,
			"abc-2": 200,
			"xyz-1": 300,
			"xyz-2": 400,
		},
	}
}

// GetAll get all key value pair
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.Database)
	if err != nil {
		http.Error(w, "error in decoding Get All", http.StatusInternalServerError)
	}
}

// GetValue get value of a key
func (h *Handler) GetValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_, ok := h.Database[params["key"]]
	if !ok {
		http.Error(w, "Key not present", http.StatusBadRequest)
		return
	}
	err := json.NewEncoder(w).Encode(h.Database[params["key"]])
	if err != nil {
		http.Error(w, "error in decoding GetValue", http.StatusInternalServerError)
	}
}

// SetValue set a key value pair
func (h *Handler) SetValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, "error in decoding SetValue", http.StatusInternalServerError)
		return
	}

	for key, value := range new {
		h.Database[key] = value
	}
	err = json.NewEncoder(w).Encode(h.Database)
	if err != nil {
		http.Error(w, "error in decoding GetValue", http.StatusInternalServerError)
	}
}

//Search a value by key value suffix
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var arr []string
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error in getting search value", http.StatusNotFound)
		return
	}
	queries := r.Form

	if len(queries["prefix"]) > 0 {
		for key := range h.Database {
			if strings.HasPrefix(key, queries["prefix"][0]) {
				arr = append(arr, key)
			}
		}
	}

	if len(queries["suffix"]) > 0 {
		for key := range h.Database {
			if strings.HasSuffix(key, queries["suffix"][0]) {
				arr = append(arr, key)
			}
		}
	}

	err = json.NewEncoder(w).Encode(arr)
	if err != nil {
		http.Error(w, "error in decoding Search query", http.StatusInternalServerError)
	}
}
