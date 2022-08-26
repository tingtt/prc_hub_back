package event

func DeleteEvent(repo EventRepository, id string) error {
	return repo.Delete(id)
}
