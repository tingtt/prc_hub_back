package event

func GetDocument(repo Repos, id string) (EventDocument, error) {
	return repo.Document.Get(id)
}
