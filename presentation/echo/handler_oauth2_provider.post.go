package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (POST /users/oauth2/{oauth_providers})
func (s Server) PostUsersOauth2OauthProviders(ctx echo.Context, oauthProviders PostUsersOauth2OauthProvidersParamsOauthProviders) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
