package models

type Teacher struct {
	UUID       string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	FirstName  string `json:"first_name" example:"Имя"`
	SecondName string `json:"second_name" example:"Фамилия"`
	MiddleName string `json:"middle_name,omitempty" example:"Отчество"`
}

type UpdateTeacherRequest struct {
	FirstName  string `json:"first_name" example:"Имя"`
	SecondName string `json:"second_name" example:"Фамилия"`
	MiddleName string `json:"middle_name,omitempty" example:"Отчество"`
}

type AddTeacherRequest struct {
	FirstName  string `json:"first_name" example:"Имя"`
	SecondName string `json:"second_name" example:"Фамилия"`
	MiddleName string `json:"middle_name,omitempty" example:"Отчество"`
}

type GetTeacherResponse struct {
	UUID     string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	FullName string `json:"full_name" example:"ФИО"`
}

type TeacherResponse struct {
	TeacherUUID string `json:"teacher_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}

type AddTeacherSchedule struct {
	TeacherUUID string `json:"teacher_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}
