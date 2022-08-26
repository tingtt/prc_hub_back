package event

import "prc_hub_back/domain/model/util"

type UpdateEventParam struct {
	Name      *string                 `json:"name"`
	Location  util.NullableJSONString `json:"location,omitempty"`
	Published *bool                   `json:"published"`
	Completed *bool                   `json:"completed"`
}

func (p UpdateEventParam) validate() error {
	if p.Name != nil {
		err := validateTitle(*p.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateEvent(repo EventRepository, id string, p UpdateEventParam) (_ Event, err error) {
	// 確認
	_, err = repo.Get(id)
	if err != nil {
		return
	}

	// バリデーション
	err = p.validate()
	if err != nil {
		return
	}

	return repo.Update(id, p)
}
