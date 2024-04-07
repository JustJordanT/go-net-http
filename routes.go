package main

import "net/http"

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/{id}", s.getUser)
	mux.HandleFunc("POST /users", s.createUser)
	// s.router.HandleFunc("/users/{id}", s.getUser())
	mux.HandleFunc("PUT /users/{id}", s.updateUser)
	mux.HandleFunc("DELETE /users/{id}", s.deleteUser)
	return mux
}
