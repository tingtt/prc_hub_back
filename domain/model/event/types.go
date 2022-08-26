package event

type Event struct {
	Id        string           `json:"id"`
	Name      string           `json:"name"`
	Location  *string          `json:"location,omitempty"`
	Documents *[]EventDocument `json:"documents,omitempty"`
	Published bool             `json:"published"`
	Completed bool             `json:"completed"`
}

type EventDocument struct {
	EventId string `json:"event_id"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}
