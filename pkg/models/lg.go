package models

type LocalGovernment struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	StateID int   `json:"state_id"`
}
