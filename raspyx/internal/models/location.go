package models

type Location struct {
	UUID string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	Name string `json:"name" example:"Автозаводская"`
}

type UpdateLocationRequest struct {
	Name string `json:"name" example:"Автозаводская"`
}

type AddLocationRequest struct {
	Name string `json:"name" example:"Автозаводская"`
}

type LocationResponse struct {
	LocationUUID string `json:"location_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}
