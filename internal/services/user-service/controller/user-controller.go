package usercontroller

import (
	"net/http"

	userservice "quups-backend/internal/services/user-service/service"
	apiutils "quups-backend/internal/utils/api"
	local_jwt "quups-backend/internal/utils/jwt"
)

func (c *controller) GetUserTeams(w http.ResponseWriter, r *http.Request) {
	response := apiutils.New(w, r)

	claims := local_jwt.GetAuthContext(r.Context())

	usrv := userservice.New(r.Context(), c.db).UserService()

	t, err := usrv.GetUserTeams(claims.Sub)
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
