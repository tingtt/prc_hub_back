package echo

import (
	"fmt"
	"prc_hub_back/domain/model/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Start(port uint, logLevel uint, gzipLevel uint, jwtIssuer string, jwtSecret string, allowOrigins []string) {
	// echoサーバーのインスタンス生成
	e := echo.New()

	// Gzipの圧縮レベル設定
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: int(gzipLevel),
	}))

	// ログレベルの設定
	e.Logger.SetLevel(log.Lvl(logLevel))

	// CORS
	if allowOrigins != nil {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: allowOrigins,
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentType,
				echo.HeaderAccept,
				echo.HeaderAuthorization,
			},
		}))
		e.Logger.Info("CORS enabled")
		e.Logger.Debugf("CORS allow origins: %v", allowOrigins)
	}

	// JWT
	jwt.InitWithSkipper(
		e,
		jwtIssuer,
		jwtSecret,
		func(c echo.Context) bool {
			// 公開エンドポイントのJWT認証をスキップ
			return c.Path() == "/users" && c.Request().Method == "POST" ||
				c.Path() == "/users/oauth2/:oauth_providers/register" && c.Request().Method == "POST" ||
				c.Path() == "/users/sign_in" && c.Request().Method == "POST" ||
				c.Path() == "/events" && c.Request().Method == "GET" ||
				c.Path() == "/events/:id" && c.Request().Method == "GET" ||
				c.Path() == "/events/:id/documents" && c.Request().Method == "GET" ||
				c.Path() == "/events/:id/documents/:document_id" && c.Request().Method == "GET" ||
				c.Path() == "/users/:id/star" && c.Request().Method == "POST" ||
				c.Path() == "/reset" && c.Request().Method == "POST"
		},
	)

	// handlerの登録
	RegisterHandlers(e, Server{})

	// echoサーバーの起動
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
