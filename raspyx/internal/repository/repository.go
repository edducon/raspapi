package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"raspyx2/internal/models"
	"raspyx2/internal/repository/groups"
	"raspyx2/internal/repository/locations"
	"raspyx2/internal/repository/rooms"
	"raspyx2/internal/repository/schedule"
	"raspyx2/internal/repository/subjectTypes"
	"raspyx2/internal/repository/subjects"
	"raspyx2/internal/repository/teachers"
)

type GroupsRepository interface {
	CreateGroup(groupData *models.Group) error

	GetAllGroups() (*[]models.Group, error)
	GetGroupsByNumber(groupNumber string) (*[]models.Group, error)
	GetGroupByNumber(groupNumber string) (*models.Group, error)
	GetGroupByUUID(groupUUID string) (*models.Group, error)

	UpdateGroup(groupUUID string, groupData *models.UpdateGroupRequest) error

	DeleteGroup(groupUUID string) error
}

type LocationsRepository interface {
	CreateLocation(locationData *models.Location) error

	GetAllLocations() (*[]models.Location, error)
	GetLocationsByName(locationName string) (*[]models.Location, error)
	GetLocationByName(locationName string) (*models.Location, error)
	GetLocationByUUID(locationUUID string) (*models.Location, error)

	UpdateLocation(locationUUID string, locationData *models.UpdateLocationRequest) error

	DeleteLocation(locationUUID string) error
}

type RoomsRepository interface {
	CreateRoom(roomData *models.Room) error

	GetAllRooms() (*[]models.Room, error)
	GetRoomsByNumber(roomNumber string) (*[]models.Room, error)
	GetRoomByNumber(roomNumber string) (*models.Room, error)
	GetRoomByUUID(roomUUID string) (*models.Room, error)

	UpdateRoom(roomUUID string, roomData *models.UpdateRoomRequest) error

	DeleteRoom(roomUUID string) error
}

type SubjectsRepository interface {
	CreateSubject(subjectData *models.Subject) error

	GetAllSubjects() (*[]models.Subject, error)
	GetSubjectsByName(subjectName string) (*[]models.Subject, error)
	GetSubjectByName(subjectName string) (*models.Subject, error)
	GetSubjectByUUID(subjectUUID string) (*models.Subject, error)

	UpdateSubject(subjectUUID string, subjectData *models.UpdateSubjectRequest) error

	DeleteSubject(subjectUUID string) error
}

type SubjectTypesRepository interface {
	CreateSubjectType(subjectTypeData *models.SubjectType) error

	GetAllSubjectTypes() (*[]models.SubjectType, error)
	GetSubjectTypesByType(subjectType string) (*[]models.SubjectType, error)
	GetSubjectTypeByType(subjectType string) (*models.SubjectType, error)
	GetSubjectTypeByUUID(subjectTypeUUID string) (*models.SubjectType, error)

	UpdateSubjectType(subjectTypeUUID string, subjectTypeData *models.UpdateSubjectTypeRequest) error

	DeleteSubjectType(subjectTypeUUID string) error
}

type TeachersRepository interface {
	CreateTeacher(teacherData *models.Teacher) error

	GetAllTeachers() (*[]models.Teacher, error)
	GetTeachersByFio(teacherFio string) (*[]models.Teacher, error)
	GetTeacherByFio(teacherFio string) (*models.Teacher, error)
	GetTeacherByUUID(teacherUUID string) (*models.Teacher, error)

	UpdateTeacher(teacherUUID string, teacherData *models.UpdateTeacherRequest) error

	DeleteTeacher(teacherUUID string) error
}

type ScheduleRepository interface {
	CreateSchedule(scheduleData *models.CreateSchedule) error

	GetScheduleByGroupUUID(groupUUID string, isSession bool) (*[]models.GetSchedule, error)
	GetScheduleByTeacherUUID(teacherUUID string, isSession bool) (*[]models.GetSchedule, error)

	GetScheduleByLocationUUID(locationUUID string, isSession bool) (*[]models.GetSchedule, error)

	UpdateSchedule(scheduleUUID string, scheduleData *models.UpdateScheduleRequest) error
	UpdateScheduleLinkByLesson(filters *models.UpdateScheduleLinkByLessonRequest) (int64, error)

	DeleteSchedule(scheduleUUID string) error
	DeleteScheduleByFilters(filters *models.DeleteScheduleFilters) error

	GetAllSchedule(isSession bool) (*[]models.GetSchedule, error)
}

type Repository struct {
	GroupsRepository
	LocationsRepository
	RoomsRepository
	SubjectsRepository
	SubjectTypesRepository
	TeachersRepository
	ScheduleRepository
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		GroupsRepository:       groups.NewGroupsRepository(pool),
		LocationsRepository:    locations.NewLocationsRepository(pool),
		RoomsRepository:        rooms.NewRoomsRepository(pool),
		SubjectsRepository:     subjects.NewSubjectsRepository(pool),
		SubjectTypesRepository: subjectTypes.NewSubjectTypesRepository(pool),
		TeachersRepository:     teachers.NewTeachersRepository(pool),
		ScheduleRepository:     schedule.NewScheduleRepository(pool),
	}
}
