package echo

import (
	"net/http"
	"prc_hub_back/application/user"

	"github.com/labstack/echo/v4"
)

func (s Server) PostUsersIdStar(ctx echo.Context, id Id) error {
	count, err := user.AddStar(uint64(id))
	if err != nil {
		return JSONMessage(ctx, http.StatusInternalServerError, err.Error())
	}
	return JSONPretty(ctx, http.StatusOK, map[string]uint64{"count": count})
}
