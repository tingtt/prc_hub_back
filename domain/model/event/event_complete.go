package event

import "prc_hub_back/domain/model/user"

func CompleteEvent(repo EventRepository, qs EventQueryService, id string, requestUser user.User) (Event, error) {
	completed := true
	return UpdateEvent(repo, qs, id, UpdateEventParam{Completed: &completed}, requestUser)
}
