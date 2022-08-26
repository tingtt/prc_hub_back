package event

func CompleteEvent(repo EventRepository, id string) (Event, error) {
	completed := true
	return UpdateEvent(repo, id, UpdateEventParam{Completed: &completed})
}
