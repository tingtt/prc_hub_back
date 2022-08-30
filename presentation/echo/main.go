package echo

import (
	"fmt"
	"prc_hub_back/domain/model/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Start(port uint, logLevel uint, gzipLevel uint, jwtIssuer string, jwtSecret string) {
	// echoサーバーのインスタンス生成
	e := echo.New()

	// Gzipの圧縮レベル設定
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: int(gzipLevel),
	}))

	// ログレベルの設定
	e.Logger.SetLevel(log.Lvl(logLevel))

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
				c.Path() == "/events/:id/documents/:document_id" && c.Request().Method == "GET"
		},
	)

	// handlerの登録
	RegisterHandlers(e, Server{})

	// echoサーバーの起動
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
