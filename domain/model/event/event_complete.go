package event

import "prc_hub_back/domain/model/user"

func CompleteEvent(repo Repos, id string, requestUser user.User) (Event, error) {
	completed := true
	return UpdateEvent(repo, id, UpdateEventParam{Completed: &completed}, requestUser)
}
