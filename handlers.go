package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the user data
	var user Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the user
	user.ID = uuid.New()

	// Save the user to the database
	result := s.db.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created user as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the Content-Type header to application/json
	var user Users
	id := r.PathValue("id")
	s.db.First(&Users{}, "id = ?", id).Scan(&user)

	json.NewEncoder(w).Encode(user)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	var user Users
	result := s.db.First(&Users{}, "id = ?", id).Scan(&user)
	if result.Error != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	s.db.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	s.db.Delete(&Users{}, id)
	w.WriteHeader(http.StatusNoContent)
}
