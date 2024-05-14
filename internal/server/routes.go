package server

import (
	"encoding/json"
	"log"
	"net/http"
	authservice "quups-backend/internal/services/auth-service/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Route("/auth", s.authController)

	// authenticated routes
	return r
}

func (s *Server) authController(r chi.Router) {
	h := authservice.New(s.repository)

	r.Post("/signin", h.Signin)
	r.Post("/signup", h.Signup)
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
