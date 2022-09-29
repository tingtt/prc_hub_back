package event

type GetDocumentQueryParam struct {
	EventId     *string `query:"event_id"`
	Name        *string `query:"name"`
	NameContain *string `query:"name_contain"`
}

func GetDocumentList(repo EventDocumentRepository, q GetDocumentQueryParam) ([]EventDocument, error) {
	// TODO: 権限によって表示を変更
	return repo.GetList(q)
}
