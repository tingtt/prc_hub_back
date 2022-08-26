package event

type GetEventQueryParam struct {
	Published       *bool   `query:"published"`
	Name            *string `query:"name"`
	NameContain     *string `query:"name_contain"`
	Location        *string `query:"location"`
	LocationContain *string `query:"location_contain"`
}

func GetEventList(repo Repos, q GetEventQueryParam) ([]Event, error) {
	return repo.Event.GetList(q)
}
