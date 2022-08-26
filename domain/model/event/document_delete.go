package event

func DeleteEventDocument(repo EventDocumentRepository, id string) error {
	return repo.Delete(id)
}
