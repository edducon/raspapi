package service

import (
	"raspyx2/config"
	"raspyx2/internal/models"
	"raspyx2/internal/repository"
	"raspyx2/internal/service/groups"
	"raspyx2/internal/service/locations"
	"raspyx2/internal/service/rooms"
	"raspyx2/internal/service/schedule"
	"raspyx2/internal/service/subjectTypes"
	"raspyx2/internal/service/subjects"
	"raspyx2/internal/service/teachers"
)

type GroupsService interface {
	CreateGroup(groupReqData *models.AddGroupRequest) (string, error)

	GetAllGroups() (*[]models.Group, error)
	GetGroupsByNumber(groupNumber string) (*[]models.Group, error)
	GetGroupByNumber(groupNumber string) (*models.Group, error)
	GetGroupByUUID(groupUUID string) (*models.Group, error)

	UpdateGroup(groupUUID string, groupData *models.UpdateGroupRequest) error

	DeleteGroup(groupUUID string) error
}

type LocationsService interface {
	CreateLocation(locationReqData *models.AddLocationRequest) (string, error)

	GetAllLocations() (*[]models.Location, error)
	GetLocationsByName(locationName string) (*[]models.Location, error)
	GetLocationByName(locationName string) (*models.Location, error)
	GetLocationByUUID(locationUUID string) (*models.Location, error)

	UpdateLocation(locationUUID string, locationData *models.UpdateLocationRequest) error

	DeleteLocation(locationUUID string) error
}

type RoomsService interface {
	CreateRoom(roomReqData *models.AddRoomRequest) (string, error)

	GetAllRooms() (*[]models.Room, error)
	GetRoomsByNumber(roomNumber string) (*[]models.Room, error)
	GetRoomByNumber(roomNumber string) (*models.Room, error)
	GetRoomByUUID(roomUUID string) (*models.Room, error)

	UpdateRoom(roomUUID string, roomData *models.UpdateRoomRequest) error

	DeleteRoom(roomUUID string) error
}

type SubjectsService interface {
	CreateSubject(subjectReqData *models.AddSubjectRequest) (string, error)

	GetAllSubjects() (*[]models.Subject, error)
	GetSubjectsByName(subjectName string) (*[]models.Subject, error)
	GetSubjectByName(subjectName string) (*models.Subject, error)
	GetSubjectByUUID(subjectUUID string) (*models.Subject, error)

	UpdateSubject(subjectUUID string, subjectData *models.UpdateSubjectRequest) error

	DeleteSubject(subjectUUID string) error
}

type SubjectTypesService interface {
	CreateSubjectType(subjectTypeReqData *models.AddSubjectTypeRequest) (string, error)

	GetAllSubjectTypes() (*[]models.SubjectType, error)
	GetSubjectTypesByType(subjectType string) (*[]models.SubjectType, error)
	GetSubjectTypeByType(subjectType string) (*models.SubjectType, error)
	GetSubjectTypeByUUID(subjectTypeUUID string) (*models.SubjectType, error)

	UpdateSubjectType(subjectTypeUUID string, subjectTypeData *models.UpdateSubjectTypeRequest) error

	DeleteSubjectType(subjectTypeUUID string) error
}

type TeachersService interface {
	CreateTeacher(teacherReqData *models.AddTeacherRequest) (string, error)

	GetAllTeachers() (*[]models.GetTeacherResponse, error)
	GetTeachersByFio(teacherFio string) (*[]models.GetTeacherResponse, error)
	GetTeacherByFio(teacherFio string) (*models.Teacher, error)
	GetTeacherByUUID(teacherUUID string) (*models.GetTeacherResponse, error)

	UpdateTeacher(teacherUUID string, teacherData *models.UpdateTeacherRequest) error

	DeleteTeacher(teacherUUID string) error
}

type ScheduleService interface {
	CreateSchedule(scheduleReqData *models.AddScheduleRequest) (string, error)

	GetScheduleByGroupNumber(groupNumber string, isSession bool) (*models.Week, error)
	GetScheduleByTeacherFio(teacherFio string, isSession bool) (*models.Week, error)

	GetScheduleByLocationName(locationName string, isSession bool) (*models.Week, error)

	UpdateSchedule(scheduleUUID string, scheduleData *models.UpdateScheduleRequest) error

	DeleteSchedule(scheduleUUID string) error
	DeleteScheduleByFilters(filters *models.DeleteScheduleFilters) error

	GetAllSchedule(isSession bool) (*[]models.GetSchedule, error)
}

type Service struct {
	GroupsService
	LocationsService
	RoomsService
	SubjectsService
	SubjectTypesService
	TeachersService
	ScheduleService
}

func NewService(cfg *config.Config, repo *repository.Repository) *Service {
	return &Service{
		GroupsService:       groups.NewGroupsService(cfg, repo),
		LocationsService:    locations.NewLocationsService(cfg, repo),
		RoomsService:        rooms.NewRoomsService(cfg, repo),
		SubjectsService:     subjects.NewSubjectsService(cfg, repo),
		SubjectTypesService: subjectTypes.NewSubjectTypesService(cfg, repo),
		TeachersService:     teachers.NewTeachersService(cfg, repo),
		ScheduleService:     schedule.NewScheduleService(cfg, repo),
	}
}
