package event

func GetEvent(repo Repos, id string) (Event, error) {
	return repo.Event.Get(id)
}
