package echo

import (
	"net/http"
	"prc_hub_back/application/user"

	"github.com/labstack/echo/v4"
)

// (GET /users)
func (s Server) GetUsers(ctx echo.Context) error {
	// Get users
	u, err := user.GetList()
	if err != nil {
		return JSONMessage(ctx, user.ErrToCode(err), err.Error())
	}
	return JSONPretty(ctx, http.StatusOK, u)
}
