package models

type EventMin struct {
	Uuid              string  `json:"uuid"`
	Name              string  `json:"name"`
	Owner             string  `json:"owner"`
	Description       string  `json:"description"`
	StartDate         int     `json:"start-date"`
	EndDate           int     `json:"end-date"`
	Rating            float32 `json:"rating"`
	MinParticipants   int     `json:"min-participants"`
	MaxParticipants   int     `json:"max-participants"`
	Place             string  `json:"place"`
	ParticipantsCount int     `json:"participants-count,omitempty"`
}
