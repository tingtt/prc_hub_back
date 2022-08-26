package main

import (
	"prc_hub_back/application/event"
	"prc_hub_back/application/oauth2"
	"prc_hub_back/application/user"
	event_inmemory "prc_hub_back/infrastructure/datasource/event/in_memory"
	oauth2_inmemory "prc_hub_back/infrastructure/datasource/oauth2/in_memory"
	user_inmemory "prc_hub_back/infrastructure/datasource/user/in_memory"
	"prc_hub_back/presentation/echo"
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
	oauth2.Init()
	oauth2.InitApplication(repositoryOAuth2)
	event.InitApplication(repositoryEvent, repositoryDocument)

	echo.Start()
}
