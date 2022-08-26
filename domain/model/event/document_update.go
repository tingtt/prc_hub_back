package event

type UpdateEventDocumentParam struct {
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

func (p UpdateEventDocumentParam) validate() error {
	if p.Name != nil {
		err := validateDocumentName(*p.Name)
		if err != nil {
			return err
		}
	}
	if p.Url != nil {
		err := validateUrl(*p.Url)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateEventDocument(repo Repos, id string, p UpdateEventDocumentParam) (_ EventDocument, err error) {
	// 確認
	_, err = repo.Document.Get(id)
	if err != nil {
		return
	}

	// バリデーション
	err = p.validate()
	if err != nil {
		return
	}

	return repo.Document.Update(id, p)
}
