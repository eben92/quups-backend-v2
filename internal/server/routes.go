package server

import (
	"encoding/json"
	"log"
	"net/http"
	authcontroller "quups-backend/internal/services/auth-service/controller"
	usercontroller "quups-backend/internal/services/user-service/controller"
	local_jwt "quups-backend/internal/utils/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.Get("/", s.HelloWorldHandler)

	// unprotected routes
	r.Route("/auth", s.authController)

	// protected routes
	r.Group(func(pr chi.Router) {
		pr.Use(local_jwt.Authenticator())

		pr.Route("/companies", s.userController)
		pr.Get("/health", s.healthHandler)
	})

	return r
}

func (s *Server) authController(r chi.Router) {
	handler := authcontroller.New(s.repository)

	r.Post("/signin", handler.Signin)
	r.Post("/signup", handler.Signup)
}

func (s *Server) userController(r chi.Router) {
	handler := usercontroller.New(s.repository)

	r.Post("/", handler.CreateCompany)
	r.Get("/", handler.GetAllCompanies)
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
	u := local_jwt.GetAuthContext(r.Context())

	log.Printf("auth ctx [%s]", u.Sub)
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
