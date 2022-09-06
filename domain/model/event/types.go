package event

type Event struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	Description *string          `json:"description"`
	Location    *string          `json:"location,omitempty"`
	Documents   *[]EventDocument `json:"documents,omitempty"`
	Published   bool             `json:"published"`
	Completed   bool             `json:"completed"`
	UserId      string           `json:"user_id"`
}

type EventDocument struct {
	EventId string `json:"event_id"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}
