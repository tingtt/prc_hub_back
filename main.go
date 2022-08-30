package main

import (
	"fmt"
	"prc_hub_back/application/event"
	"prc_hub_back/application/oauth2"
	"prc_hub_back/application/user"
	"prc_hub_back/domain/model/flag_with_env"
	event_inmemory "prc_hub_back/infrastructure/datasource/event/in_memory"
	oauth2_inmemory "prc_hub_back/infrastructure/datasource/oauth2/in_memory"
	user_inmemory "prc_hub_back/infrastructure/datasource/user/in_memory"
	"prc_hub_back/presentation/echo"
)

// コマンドライン引数 / 環境変数
var (
	adminEmail    = flag_with_env.String("admin-email", "ADMIN_EMAIL", "", "Admin user email")
	adminPassword = flag_with_env.String("admin-password", "ADMIN_PASSWORD", "", "Admin user password")
)

var (
	repositoryUser     = user_inmemory.Repository{}
	repositoryOAuth2   = oauth2_inmemory.Repository{}
	repositoryEvent    = event_inmemory.RepositoryEvent{}
	repositoryDocument = event_inmemory.RepositoryEventDocument{}
)

func main() {
	// Init application services
	user.InitApplication(repositoryUser)
	oauth2.InitApplication(repositoryOAuth2)
	event.InitApplication(repositoryEvent, repositoryDocument)

	// Migrate admin user
	fmt.Printf("adminEmail: %v\n", *adminEmail)
	fmt.Printf("adminPassword: %v\n", *adminPassword)
	err := user.SaveAdmin(*adminEmail, *adminPassword)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	users, _ := user.GetList(user.GetUserListQuery{})
	fmt.Printf("users: %v\n", users)

	echo.Start()
}
