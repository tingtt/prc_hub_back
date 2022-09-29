package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEvent) UpdateDocument(id string, p event.UpdateEventDocumentParam) (_ event.EventDocument, err error) {
	ed, err := r.GetDocument(id)
	if err != nil {
		return
	}

	if p.Name != nil {
		ed.Name = *p.Name
	}
	if p.Url != nil {
		ed.Url = *p.Url
	}

	err = r.DeleteDocument(id)
	if err != nil {
		return
	}
	_, err = r.AddDocument(ed)
	if err != nil {
		return
	}

	return ed, nil
}
