package models

type SubjectType struct {
	UUID string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	Type string `json:"type" example:"Лекция"`
}

type UpdateSubjectTypeRequest struct {
	Type string `json:"type" example:"Лекция"`
}

type AddSubjectTypeRequest struct {
	Type string `json:"type" example:"Лекция"`
}

type SubjectTypeResponse struct {
	SubjectTypeUUID string `json:"subject_type_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}
