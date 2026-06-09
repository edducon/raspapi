package models

type User struct {
	UUID         string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	Username     string `json:"username" example:"username"`
	PasswordHash string `json:"password" example:"password"`
	AccessLevel  int    `json:"access_level" example:"100"`
}

type GetUserResponse struct {
	UUID        string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	Username    string `json:"username" example:"username"`
	AccessLevel int    `json:"access_level" example:"100"`
}

type UpdateUserRequest struct {
	Username    string `json:"username" example:"username"`
	AccessLevel int    `json:"access_level" example:"50"`
}

type AddUserRequest struct {
	UUID        string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
	Username    string `json:"username" example:"username"`
	Password    string `json:"password" example:"password"`
	AccessLevel int    `json:"access_level" example:"50"`
}

type UserResponse struct {
	UserUUID string `json:"uuid" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}
