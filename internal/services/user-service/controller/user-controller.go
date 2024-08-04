package usercontroller

import (
	"net/http"

	userservice "quups-backend/internal/services/user-service/service"
	apiutils "quups-backend/internal/utils/api"

	"github.com/go-chi/chi/v5"
)

func (c *controller) GetUserCompanies(w http.ResponseWriter, r *http.Request) {
	response := apiutils.New(w, r)
	usrv := userservice.NewUserService(r.Context(), c.db)

	t, err := usrv.GetUserTeams()

	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusOK,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}
	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success",
		Results:    t,
	})
}

func (c *controller) GetUserCompany(w http.ResponseWriter, r *http.Request) {
	response := apiutils.New(w, r)
	usrv := userservice.NewUserService(r.Context(), c.db)

	companyID := chi.URLParam(r, "id")

	t, err := usrv.GetUserTeam(companyID)

	if err != nil {

		response.WrapInApiResponse(&apiutils.ApiResponseParams{
			StatusCode: http.StatusOK,
			Message:    err.Error(),
			Results:    nil,
		})

		return
	}
	response.WrapInApiResponse(&apiutils.ApiResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success",
		Results:    t,
	})
}
