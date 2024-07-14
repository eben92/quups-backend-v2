package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	authcontroller "quups-backend/internal/services/auth-service/controller"
	paymentcontroller "quups-backend/internal/services/payment-service/controller"
	usercontroller "quups-backend/internal/services/user-service/controller"
	local_jwt "quups-backend/internal/utils/jwt"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4173", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Accept-Encoding"},
		AllowCredentials: true,
	}))

	v1 := chi.NewRouter()

	v1.Get("/", s.HelloWorldHandler)
	// unprotected routes
	v1.Route("/auth", s.authController)
	// protected routes
	v1.Group(func(pr chi.Router) {
		pr.Use(local_jwt.Authenticator())
		pr.Route("/user", s.userController)
		pr.Route("/companies", s.companyController)
		pr.Route("/payments", s.paymentController)
	})

	v1.Get("/health", s.healthHandler)

	r.Mount("/api/v1", v1)
	return r
}

func (s *Server) authController(r chi.Router) {
	handler := authcontroller.New(s.db)

	r.Post("/signin", handler.Signin)
	r.Post("/signup", handler.Signup)
	r.Post("/signout", handler.Signout)
}

func (s *Server) companyController(r chi.Router) {
	handler := usercontroller.NewCompanyController(s.db)

	r.Post("/", handler.CreateCompany)
	r.Get("/", handler.GetAllCompanies)
	r.Get("/{id}", handler.GetCompanyByID)
	r.Get("/name/{name}", handler.GetCompanyByName)
	r.Get("/name/exists", handler.GetCompanyNameAvailability)
}

// payment controller
func (s *Server) paymentController(r chi.Router) {
	handler := paymentcontroller.NewPaymentController(s.db)

	r.Get("/supported-banks", handler.GetBankList)
}

func (s *Server) userController(r chi.Router) {
	handler := usercontroller.NewUserController(s.db)
	authhandler := authcontroller.New(s.db)

	r.Get("/teams", handler.GetUserTeams)
	r.Post("/account", authhandler.AccountSignin)
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
