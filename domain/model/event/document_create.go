package event

import "prc_hub_back/domain/model/util"

type CreateEventDocumentParam struct {
	EventId string `json:"event_id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

func (p CreateEventDocumentParam) validate(repo EventRepository) error {
	err := validateDocumentName(p.Name)
	if err != nil {
		return err
	}
	err = validateUrl(p.Url)
	if err != nil {
		return err
	}
	err = validateEventId(repo, p.EventId)
	if err != nil {
		return err
	}
	return nil
}

func CreateEventDocument(repo Repos, p CreateEventDocumentParam) (_ EventDocument, err error) {
	err = p.validate(repo.Event)
	if err != nil {
		return
	}

	return repo.Document.Add(EventDocument{
		Id:      util.UUID(),
		EventId: p.EventId,
		Name:    p.Name,
		Url:     p.Url,
	})
}
