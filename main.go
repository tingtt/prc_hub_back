package main

import (
	"fmt"
	"os"
	"prc_hub_back/application/event"
	"prc_hub_back/application/user"
	"prc_hub_back/domain/model/flag_with_env"
	"prc_hub_back/presentation/echo"
)

// コマンドライン引数 / 環境変数
var (
	port          = flag_with_env.Uint("port", "PORT", 1323, "Server port")
	logLevel      = flag_with_env.Uint("log-level", "LOG_LEVEL", 2, "Log level (1: 'DEBUG', 2: 'INFO', 3: 'WARN', 4: 'ERROR', 5: 'OFF', 6: 'PANIC', 7: 'FATAL'")
	gzipLevel     = flag_with_env.Uint("gzip-level", "GZIP_LEVEL", 6, "Gzip compression level")
	issuer        = flag_with_env.String("jwt-issuer", "JWT_ISSUER", "prc_hub-api", "JWT issuer")
	secret        = flag_with_env.String("jwt-secret", "JWT_SECRET", "", "JWT secret")
	adminEmail    = flag_with_env.String("admin-email", "ADMIN_EMAIL", "", "Admin user email")
	adminPassword = flag_with_env.String("admin-password", "ADMIN_PASSWORD", "", "Admin user password")
	allowOrigins  = flag_with_env.Array("allow-origin", "CORS allow origins")

	mysqlHost     = flag_with_env.String("mysql-host", "MYSQL_HOST", "localhost", "")
	mysqlPort     = flag_with_env.Uint("mysql-port", "MYSQL_PORT", 3306, "")
	mysqlDB       = flag_with_env.String("mysql-db", "MYSQL_DATABASE", "prc_hub", "")
	mysqlUser     = flag_with_env.String("mysql-user", "MYSQL_USER", "prc_hub", "")
	mysqlPassword = flag_with_env.String("mysql-password", "MYSQL_PASSWORD", "", "")
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

	// Init application services
	user.Init(*mysqlUser, *mysqlPassword, *mysqlHost, *mysqlPort, *mysqlDB)
	event.Init(*mysqlUser, *mysqlPassword, *mysqlHost, *mysqlPort, *mysqlDB)

	// Migrate admin user
	fmt.Printf("adminEmail: %v\n", *adminEmail)
	err := user.SaveAdmin(*adminEmail, *adminPassword)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	echo.Start(*port, *logLevel, *gzipLevel, *issuer, *secret, *allowOrigins)
}
