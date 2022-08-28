package event

import "prc_hub_back/domain/model/user"

type GetEventQueryParam struct {
	Published       *bool   `query:"published"`
	Name            *string `query:"name"`
	NameContain     *string `query:"name_contain"`
	Location        *string `query:"location"`
	LocationContain *string `query:"location_contain"`
}

func GetEventList(repo Repos, q GetEventQueryParam, requestUser user.User) ([]Event, error) {
	// TODO: 権限によって表示を変更
	return repo.Event.GetList(q)
}
