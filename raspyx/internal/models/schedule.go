package models

import "time"

type Schedule struct {
	UUID            string     `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	GroupUUID       string     `json:"group_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	SubjectUUID     string     `json:"subject_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	SubjectTypeUUID string     `json:"subject_type_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	LocationUUID    string     `json:"location_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	TeachersUUID    string     `json:"teachers_uuid"`
	RoomsUUID       string     `json:"rooms_uuid"`
	StartTime       *time.Time `json:"start_time" example:"09:00:00"`
	EndTime         *time.Time `json:"end_time" example:"10:30:00"`
	StartDate       *time.Time `json:"start_date" example:"2025-09-01"`
	EndDate         *time.Time `json:"end_date" example:"2025-12-31"`
	Weekday         int        `json:"weekday" example:"1"`
	Link            string     `json:"link" example:"https://rasp.dmami.ru/"`
	IsSession       bool       `json:"is_session" example:"false"`
}

type CreateSchedule struct {
	UUID            string               `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	GroupUUID       string               `json:"group_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	SubjectUUID     string               `json:"subject_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	SubjectTypeUUID string               `json:"subject_type_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	LocationUUID    string               `json:"location_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	TeachersUUID    []AddTeacherSchedule `json:"teachers_uuid"`
	RoomsUUID       []AddRoomSchedule    `json:"rooms_uuid"`
	StartTime       string               `json:"start_time" example:"09:00:00"`
	EndTime         string               `json:"end_time" example:"10:30:00"`
	StartDate       string               `json:"start_date" example:"2025-09-01"`
	EndDate         string               `json:"end_date" example:"2025-12-31"`
	Weekday         int                  `json:"weekday" example:"1"`
	Link            string               `json:"link" example:"https://rasp.dmami.ru/"`
	IsSession       bool                 `json:"is_session" example:"false"`
}
type AddScheduleRequest struct {
	GroupUUID       string               `json:"group_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	SubjectUUID     string               `json:"subject_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	SubjectTypeUUID string               `json:"subject_type_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	LocationUUID    string               `json:"location_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	TeachersUUID    []AddTeacherSchedule `json:"teachers_uuid"`
	RoomsUUID       []AddRoomSchedule    `json:"rooms_uuid"`
	StartTime       string               `json:"start_time" example:"09:00:00"`
	EndTime         string               `json:"end_time" example:"10:30:00"`
	StartDate       string               `json:"start_date" example:"2025-09-01"`
	EndDate         string               `json:"end_date" example:"2025-12-31"`
	Weekday         int                  `json:"weekday" example:"1"`
	Link            string               `json:"link" example:"https://rasp.dmami.ru/"`
	IsSession       bool                 `json:"is_session" example:"false"`
}

type UpdateScheduleRequest struct {
	GroupUUID       string               `json:"group_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	SubjectUUID     string               `json:"subject_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	SubjectTypeUUID string               `json:"subject_type_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	LocationUUID    string               `json:"location_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	TeachersUUID    []AddTeacherSchedule `json:"teachers_uuid"`
	RoomsUUID       []AddRoomSchedule    `json:"rooms_uuid"`
	StartTime       string               `json:"start_time" example:"09:00:00"`
	EndTime         string               `json:"end_time" example:"10:30:00"`
	StartDate       string               `json:"start_date" example:"2025-09-01"`
	EndDate         string               `json:"end_date" example:"2025-12-31"`
	Weekday         int                  `json:"weekday" example:"1"`
	Link            string               `json:"link" example:"https://rasp.dmami.ru/"`
	IsSession       bool                 `json:"is_session" example:"false"`
}

type UpdateScheduleLinkByLessonRequest struct {
	GroupUUID   string
	SubjectName string
	SubjectType string
	StartTime   string
	StartDate   string
	EndDate     string
	Weekday     int
	Link        string
	IsSession   bool
}

type ScheduleResponse struct {
	ScheduleUUID string `json:"schedule_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}

type Week struct {
	Monday    *Day `json:"monday"`
	Tuesday   *Day `json:"tuesday"`
	Wednesday *Day `json:"wednesday"`
	Thursday  *Day `json:"thursday"`
	Friday    *Day `json:"friday"`
	Saturday  *Day `json:"saturday"`
}

type Day struct {
	First   *[]Pair `json:"1"`
	Second  *[]Pair `json:"2"`
	Third   *[]Pair `json:"3"`
	Fourth  *[]Pair `json:"4"`
	Fifth   *[]Pair `json:"5"`
	Sixth   *[]Pair `json:"6"`
	Seventh *[]Pair `json:"7"`
}

type Pair struct {
	Group       *Group                `json:"group,omitempty"`
	Subject     *Subject              `json:"subject"`
	SubjectType *SubjectType          `json:"subject_type"`
	Location    *Location             `json:"location"`
	Teachers    *[]GetTeacherResponse `json:"teachers,omitempty"`
	Rooms       *[]Room               `json:"rooms,omitempty"`
	StartDate   string                `json:"start_date" example:"2025-02-01"`
	EndDate     string                `json:"end_date" example:"2025-06-01"`
	Link        string                `json:"link,omitempty" example:"https://online.mospolytech.ru/"`
}

type GetSchedule struct {
	Group       *Group                `json:"group,omitempty"`
	Subject     *Subject              `json:"subject"`
	SubjectType *SubjectType          `json:"subject_type"`
	Location    *Location             `json:"location"`
	Teachers    *[]GetTeacherResponse `json:"teachers"`
	Rooms       *[]Room               `json:"rooms,omitempty"`
	StartTime   *time.Time            `json:"start_time"`
	EndTime     *time.Time            `json:"end_time"`
	StartDate   *time.Time            `json:"start_date"`
	EndDate     *time.Time            `json:"end_date"`
	Weekday     int                   `json:"weekday"`
	Link        string                `json:"link,omitempty"`
	IsSession   bool                  `json:"is_session"`
}

type DeleteScheduleFilters struct {
	GroupUUID   string
	StartTime   string
	EndTime     string
	StartDate   string
	EndDate     string
	SubjectUUID string
	Weekday     *int
	IsSession   *bool
}

type GetScheduleByGroupNumberParams struct {
	IsSession bool `form:"is_session" binding:"omitempty"`
}

type GetScheduleByTeacherFioParams struct {
	IsSession bool `form:"is_session" binding:"omitempty"`
}
