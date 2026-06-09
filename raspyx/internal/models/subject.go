package models

type Subject struct {
	UUID string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	Name string `json:"name" example:"Иностранный язык"`
}

type UpdateSubjectRequest struct {
	Name string `json:"name" example:"Иностранный язык"`
}

type AddSubjectRequest struct {
	Name string `json:"name" example:"Иностранный язык"`
}

type SubjectResponse struct {
	SubjectUUID string `json:"subject_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}
