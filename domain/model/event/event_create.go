package event

import "prc_hub_back/domain/model/util"

type CreateEventParam struct {
	Name      string  `json:"name"`
	Location  *string `json:"location,omitempty"`
	Published bool    `json:"published"`
	Completed bool    `json:"completed"`
}

func (p CreateEventParam) validate() error {
	err := validateTitle(p.Name)
	if err != nil {
		return err
	}
	return nil
}

func CreateEvent(repo EventRepository, p CreateEventParam) (_ Event, err error) {
	// バリデーション
	err = p.validate()
	if err != nil {
		return
	}

	return repo.Add(Event{
		Id:        util.UUID(),
		Name:      p.Name,
		Location:  p.Location,
		Published: p.Published,
		Completed: p.Completed,
	})
}
