package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (DELETE /users/oauth2/{oauth_providers})
func (s Server) DeleteUsersOauth2OauthProviders(ctx echo.Context, oauthProviders DeleteUsersOauth2OauthProvidersParamsOauthProviders) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
