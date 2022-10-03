package main

import (
	"fmt"
	"os"
	"prc_hub_back/application/event"
	"prc_hub_back/application/oauth2"
	"prc_hub_back/application/user"
	"prc_hub_back/application/webhook"
	"prc_hub_back/domain/model/flag_with_env"
	"prc_hub_back/domain/model/webhook_line_notify"
	event_mysql "prc_hub_back/infrastructure/datasource/event/mysql"
	oauth2_mysql "prc_hub_back/infrastructure/datasource/oauth2/mysql"
	user_mysql "prc_hub_back/infrastructure/datasource/user/mysql"
	"prc_hub_back/presentation/echo"
)

// コマンドライン引数 / 環境変数
var (
	githubClientId     = flag_with_env.String("github-clint-id", "OAUTH2_CLIENT_ID_GITHUB", "", "OAuth2 client-id for github,com")
	githubClientSecret = flag_with_env.String("github-clint-secret", "OAUTH2_CLIENT_SECRET_GITHUB", "", "OAuth2 client-secret for github,com")
	port               = flag_with_env.Uint("port", "PORT", 1323, "Server port")
	logLevel           = flag_with_env.Uint("log-level", "LOG_LEVEL", 2, "Log level (1: 'DEBUG', 2: 'INFO', 3: 'WARN', 4: 'ERROR', 5: 'OFF', 6: 'PANIC', 7: 'FATAL'")
	gzipLevel          = flag_with_env.Uint("gzip-level", "GZIP_LEVEL", 6, "Gzip compression level")
	issuer             = flag_with_env.String("jwt-issuer", "JWT_ISSUER", "prc_hub-api", "JWT issuer")
	secret             = flag_with_env.String("jwt-secret", "JWT_SECRET", "", "JWT secret")
	adminEmail         = flag_with_env.String("admin-email", "ADMIN_EMAIL", "", "Admin user email")
	adminPassword      = flag_with_env.String("admin-password", "ADMIN_PASSWORD", "", "Admin user password")
	lineNotifyToken    = flag_with_env.String("line-notify-token", "LINE_NOTIFY_TOKEN", "", "LINE notify token")
	frontEndUrl        = flag_with_env.String("frontend-url", "FRONTEND_URL", "", "Frontend url")
	allowOrigins       = flag_with_env.Array("allow-origin", "CORS allow origins")

	mysqlHost     = flag_with_env.String("mysql-host", "MYSQL_HOST", "localhost", "")
	mysqlPort     = flag_with_env.Uint("mysql-port", "MYSQL_PORT", 3306, "")
	mysqlDB       = flag_with_env.String("mysql-db", "MYSQL_DATABASE", "prc_hub", "")
	mysqlUser     = flag_with_env.String("mysql-user", "MYSQL_USER", "prc_hub", "")
	mysqlPassword = flag_with_env.String("mysql-password", "MYSQL_PASSWORD", "", "")
)

var (
	repositoryUser            = user_mysql.Repository{}
	repositoryOAuth2          = oauth2_mysql.Repository{}
	repositoryEvent           = event_mysql.RepositoryEvent{}
	queryServiceEvent         = event_mysql.QueryServiceEvent{}
	webhookProviderLineNotify = webhook_line_notify.WebHookLineNotify{}
)

func main() {
	// コマンドライン引数 / 環境変数 の取得
	flag_with_env.Parse()
	if *issuer == "" {
		fmt.Println("`--jwt-issuer` option is required")
		os.Exit(1)
	}
	if *secret == "" {
		fmt.Println("`--jwt-secret` option is required")
		os.Exit(1)
	}
	if *adminEmail == "" {
		fmt.Println("`--admin-email` option is required")
		os.Exit(1)
	}
	if *adminPassword == "" {
		fmt.Println("`--admin-password` option is required")
		os.Exit(1)
	}
	if *frontEndUrl == "" {
		fmt.Println("`--frontend-url` option is required")
		os.Exit(1)
	}

	// Init repository
	user_mysql.InitRepository(*mysqlUser, *mysqlPassword, *mysqlHost, *mysqlPort, *mysqlDB)
	event_mysql.InitRepository(*mysqlUser, *mysqlPassword, *mysqlHost, *mysqlPort, *mysqlDB)
	oauth2_mysql.InitRepository(*mysqlUser, *mysqlPassword, *mysqlHost, *mysqlPort, *mysqlDB)

	// Init application services
	user.InitApplication(repositoryUser)
	oauth2.InitApplication(repositoryOAuth2, *githubClientId, *githubClientSecret)
	event.InitApplication(repositoryEvent, queryServiceEvent)
	webhook.InitApplication(
		*frontEndUrl,
		webhook.Provider{
			Provider: webhookProviderLineNotify,
			Token:    *lineNotifyToken,
		},
	)

	// Migrate admin user
	fmt.Printf("adminEmail: %v\n", *adminEmail)
	err := user.SaveAdmin(*adminEmail, *adminPassword)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	echo.Start(*port, *logLevel, *gzipLevel, *issuer, *secret, *allowOrigins)
}
