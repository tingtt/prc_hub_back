package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEvent) GetDocument(id string) (_ event.EventDocument, err error) {
	for _, d := range dataDocument {
		if d.Id == id {
			return d, nil
		}
	}
	err = event.ErrEventDocumentNotFound
	return
}
