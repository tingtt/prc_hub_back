package user_inmemory

import "prc_hub_back/domain/model/event"

func (r Repository) Delete(id string) (err error) {
	old := data
	data = nil
	err = event.ErrEventNotFound
	for _, d := range old {
		if d.Id == id {
			err = nil
		} else {
			data = append(data, d)
		}
	}
	return
}
