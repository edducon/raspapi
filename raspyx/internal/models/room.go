package models

type Room struct {
	UUID   string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	Number string `json:"number" example:"ав4810"`
}

type UpdateRoomRequest struct {
	Number string `json:"number" example:"ав4810"`
}

type AddRoomRequest struct {
	Number string `json:"number" example:"ав4810"`
}

type RoomResponse struct {
	RoomUUID string `json:"room_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}

type AddRoomSchedule struct {
	RoomUUID string `json:"room_uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}
