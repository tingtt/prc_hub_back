package echo

import (
	"fmt"
	"prc_hub_back/domain/model/flag_with_env"
	"prc_hub_back/domain/model/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// コマンドライン引数 / 環境変数
var (
	port      = flag_with_env.Uint("port", "PORT", 1323, "Server port")
	logLevel  = flag_with_env.Uint("log-level", "LOG_LEVEL", 2, "Log level (1: 'DEBUG', 2: 'INFO', 3: 'WARN', 4: 'ERROR', 5: 'OFF', 6: 'PANIC', 7: 'FATAL'")
	gzipLevel = flag_with_env.Uint("gzip-level", "GZIP_LEVEL", 6, "Gzip compression level")
)

func Start() {
	// コマンドライン引数 / 環境変数 の取得
	flag_with_env.Parse()

	// echoサーバーのインスタンス生成
	e := echo.New()

	// Gzipの圧縮レベル設定
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: int(*gzipLevel),
	}))

	// ログレベルの設定
	e.Logger.SetLevel(log.Lvl(*logLevel))

	// JWT
	jwt.InitWithSkipper(
		e,
		func(c echo.Context) bool {
			// 公開エンドポイントのJWT認証をスキップ
			return c.Path() == "/users" && c.Request().Method == "POST" ||
				c.Path() == "/users/oauth2/:oauth_providers/register" && c.Request().Method == "POST" ||
				c.Path() == "/users/sign_in" && c.Request().Method == "POST" ||
				c.Path() == "/events" && c.Request().Method == "GET"
		},
	)

	// handlerの登録
	RegisterHandlers(e, Server{})

	// echoサーバーの起動
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
