package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"raspyx2/config"
	"raspyx2/internal/handler/groups"
	"raspyx2/internal/handler/locations"
	"raspyx2/internal/handler/rooms"
	"raspyx2/internal/handler/schedule"
	"raspyx2/internal/handler/subjectTypes"
	"raspyx2/internal/handler/subjects"
	"raspyx2/internal/handler/teachers"
	"raspyx2/internal/service"
)

type GroupsHandler interface {
	CreateGroup(ctx *gin.Context)

	GetAllGroups(ctx *gin.Context)
	GetGroupsByNumber(ctx *gin.Context)
	GetGroupByUUID(ctx *gin.Context)

	UpdateGroup(ctx *gin.Context)

	DeleteGroup(ctx *gin.Context)
}

type LocationsHandler interface {
	CreateLocation(ctx *gin.Context)

	GetAllLocations(ctx *gin.Context)
	GetLocationsByName(ctx *gin.Context)
	GetLocationByUUID(ctx *gin.Context)

	UpdateLocation(ctx *gin.Context)

	DeleteLocation(ctx *gin.Context)
}

type RoomsHandler interface {
	CreateRoom(ctx *gin.Context)

	GetAllRooms(ctx *gin.Context)
	GetRoomsByNumber(ctx *gin.Context)
	GetRoomByUUID(ctx *gin.Context)

	UpdateRoom(ctx *gin.Context)

	DeleteRoom(ctx *gin.Context)
}

type SubjectsHandler interface {
	CreateSubject(ctx *gin.Context)

	GetAllSubjects(ctx *gin.Context)
	GetSubjectsByName(ctx *gin.Context)
	GetSubjectByUUID(ctx *gin.Context)

	UpdateSubject(ctx *gin.Context)

	DeleteSubject(ctx *gin.Context)
}

type SubjectTypesHandler interface {
	CreateSubjectType(ctx *gin.Context)

	GetAllSubjectTypes(ctx *gin.Context)
	GetSubjectTypesByType(ctx *gin.Context)
	GetSubjectTypeByUUID(ctx *gin.Context)

	UpdateSubjectType(ctx *gin.Context)

	DeleteSubjectType(ctx *gin.Context)
}

type TeachersHandler interface {
	CreateTeacher(ctx *gin.Context)

	GetAllTeachers(ctx *gin.Context)
	GetTeachersByFio(ctx *gin.Context)
	GetTeacherByUUID(ctx *gin.Context)

	UpdateTeacher(ctx *gin.Context)

	DeleteTeacher(ctx *gin.Context)
}

type ScheduleHandler interface {
	CreateSchedule(ctx *gin.Context)

	GetScheduleByGroupNumber(ctx *gin.Context)
	GetScheduleByTeacherFio(ctx *gin.Context)

	GetScheduleByLocationName(ctx *gin.Context)

	UpdateSchedule(ctx *gin.Context)

	DeleteSchedule(ctx *gin.Context)

	GetAllSchedule(ctx *gin.Context)
}

type Handler struct {
	log      *slog.Logger
	cfg      *config.Config
	services *service.Service
	GroupsHandler
	LocationsHandler
	RoomsHandler
	SubjectsHandler
	SubjectTypesHandler
	TeachersHandler
	ScheduleHandler
}

func NewHandler(log *slog.Logger, cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		log:                 log,
		cfg:                 cfg,
		GroupsHandler:       groups.NewGroupsHandler(log, cfg, services),
		LocationsHandler:    locations.NewLocationsHandler(log, cfg, services),
		RoomsHandler:        rooms.NewRoomsHandler(log, cfg, services),
		SubjectsHandler:     subjects.NewSubjectsHandler(log, cfg, services),
		SubjectTypesHandler: subjectTypes.NewSubjectTypesHandler(log, cfg, services),
		TeachersHandler:     teachers.NewTeachersHandler(log, cfg, services),
		ScheduleHandler:     schedule.NewScheduleHandler(log, cfg, services),
	}
}
