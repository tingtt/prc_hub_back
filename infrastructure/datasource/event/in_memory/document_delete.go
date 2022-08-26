package event_inmemory

import (
	"prc_hub_back/domain/model/event"
)

func (r RepositoryEventDocument) Delete(id string) (err error) {
	old := dataDocument
	dataDocument = nil
	err = event.ErrEventDocumentNotFound
	for _, d := range old {
		if d.Id == id {
			err = nil
		} else {
			dataDocument = append(dataDocument, d)
		}
	}
	return
}
