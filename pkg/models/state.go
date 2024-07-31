package models

type State struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CountryID int    `json:"country_id"`
}
