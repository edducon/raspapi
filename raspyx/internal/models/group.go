package models

type Group struct {
	UUID   string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	Number string `json:"number" example:"221-352"`
}

type UpdateGroupRequest struct {
	Number string `json:"number" example:"221-352"`
}

type AddGroupRequest struct {
	Number string `json:"number" example:"221-352"`
}

type GroupResponse struct {
	GroupUUID string `json:"group_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}
