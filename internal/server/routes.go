package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	authcontroller "quups-backend/internal/services/auth-service/controller"
	usercontroller "quups-backend/internal/services/user-service/controller"
	local_jwt "quups-backend/internal/utils/jwt"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://vendor.quups.app"},
		AllowedOrigins:   []string{"http://localhost:4173", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Accept-Encoding"},
		AllowCredentials: true,
	}))

	r.Get("/", s.HelloWorldHandler)

	// unprotected routes
	r.Route("/auth", s.authController)

	// protected routes
	r.Group(func(pr chi.Router) {
		pr.Use(local_jwt.Authenticator())
		pr.Route("/user", s.userController)
		pr.Route("/companies", s.companyController)
		pr.Route("/payments", s.paymentController)
	})

	r.Get("/health", s.healthHandler)
	return r
}

func (s *Server) authController(r chi.Router) {
	handler := authcontroller.New(s.db)

	r.Post("/signin", handler.Signin)
	r.Post("/signup", handler.Signup)
}

func (s *Server) companyController(r chi.Router) {
	handler := usercontroller.NewCompanyController(s.db)

	r.Post("/", handler.CreateCompany)
	r.Get("/", handler.GetAllCompanies)
	r.Get("/{id}", handler.GetCompanyByID)
	r.Get("/name/{name}", handler.GetCompanyByName)
}

// payment controller
func (s *Server) paymentController(r chi.Router) {
	handler := usercontroller.NewPaymentController(s.db)

	r.Get("/supported-banks", handler.GetBankList)
}

func (s *Server) userController(r chi.Router) {
	handler := usercontroller.NewUserController(s.db)

	r.Get("/teams", handler.GetUserTeams)
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
	// u := local_jwt.GetAuthContext(r.Context())

	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
