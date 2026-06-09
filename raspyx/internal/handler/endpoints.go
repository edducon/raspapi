package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "raspyx2/docs"
)

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

	raspyx := routes.Group("/raspyx")

	raspyx.Use(h.RequestLogger())

	raspyx.StaticFS("/docs", http.Dir("./docs"))

	api := raspyx.Group("/api/v2")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		groups := api.Group("/groups").Use(h.Auth()).Use(h.CheckAL(99))
		{
			groups.POST("/", h.GroupsHandler.CreateGroup)
			groups.GET("/", h.GroupsHandler.GetAllGroups)
			groups.GET("/number/:group_number", h.GroupsHandler.GetGroupsByNumber)
			groups.GET("/uuid/:group_uuid", h.GroupsHandler.GetGroupByUUID)
			groups.PUT("/:group_uuid", h.GroupsHandler.UpdateGroup)
			groups.DELETE("/:group_uuid", h.GroupsHandler.DeleteGroup)
		}

		locations := api.Group("/locations").Use(h.Auth()).Use(h.CheckAL(99))
		{
			locations.POST("/", h.LocationsHandler.CreateLocation)
			locations.GET("/", h.LocationsHandler.GetAllLocations)
			locations.GET("/number/:location_name", h.LocationsHandler.GetLocationsByName)
			locations.GET("/uuid/:location_uuid", h.LocationsHandler.GetLocationByUUID)
			locations.PUT("/:location_uuid", h.LocationsHandler.UpdateLocation)
			locations.DELETE("/:location_uuid", h.LocationsHandler.DeleteLocation)
		}

		rooms := api.Group("/rooms").Use(h.Auth()).Use(h.CheckAL(99))
		{
			rooms.POST("/", h.RoomsHandler.CreateRoom)
			rooms.GET("/", h.RoomsHandler.GetAllRooms)
			rooms.GET("/number/:room_number", h.RoomsHandler.GetRoomsByNumber)
			rooms.GET("/uuid/:room_uuid", h.RoomsHandler.GetRoomByUUID)
			rooms.PUT("/:room_uuid", h.RoomsHandler.UpdateRoom)
			rooms.DELETE("/:room_uuid", h.RoomsHandler.DeleteRoom)
		}

		subjects := api.Group("/subjects").Use(h.Auth()).Use(h.CheckAL(99))
		{
			subjects.POST("/", h.SubjectsHandler.CreateSubject)
			subjects.GET("/", h.SubjectsHandler.GetAllSubjects)
			subjects.GET("/name/:subject_name", h.SubjectsHandler.GetSubjectsByName)
			subjects.GET("/uuid/:subject_uuid", h.SubjectsHandler.GetSubjectByUUID)
			subjects.PUT("/:subject_uuid", h.SubjectsHandler.UpdateSubject)
			subjects.DELETE("/:subject_uuid", h.SubjectsHandler.DeleteSubject)
		}

		subjectTypes := api.Group("/subject_types").Use(h.Auth()).Use(h.CheckAL(99))
		{
			subjectTypes.POST("/", h.SubjectTypesHandler.CreateSubjectType)
			subjectTypes.GET("/", h.SubjectTypesHandler.GetAllSubjectTypes)
			subjectTypes.GET("/type/:subject_type", h.SubjectTypesHandler.GetSubjectTypesByType)
			subjectTypes.GET("/uuid/:subject_type_uuid", h.SubjectTypesHandler.GetSubjectTypeByUUID)
			subjectTypes.PUT("/:subject_type_uuid", h.SubjectTypesHandler.UpdateSubjectType)
			subjectTypes.DELETE("/:subject_type_uuid", h.SubjectTypesHandler.DeleteSubjectType)
		}

		teachers := api.Group("/teachers").Use(h.Auth()).Use(h.CheckAL(99))
		{
			teachers.POST("/", h.TeachersHandler.CreateTeacher)
			teachers.GET("/", h.TeachersHandler.GetAllTeachers)
			teachers.GET("/fio/:teacher_fio", h.TeachersHandler.GetTeachersByFio)
			teachers.GET("/uuid/:teacher_uuid", h.TeachersHandler.GetTeacherByUUID)
			teachers.PUT("/:teacher_uuid", h.TeachersHandler.UpdateTeacher)
			teachers.DELETE("/:teacher_uuid", h.TeachersHandler.DeleteTeacher)
		}

		scheduleAuth := api.Group("/schedule").Use(h.Auth()).Use(h.CheckAL(99))
		schedule := api.Group("/schedule")
		{
			scheduleAuth.POST("/", h.ScheduleHandler.CreateSchedule)
			schedule.GET("/all", h.ScheduleHandler.GetAllSchedule)
			schedule.GET("/group_number/:group_number", h.ScheduleHandler.GetScheduleByGroupNumber)
			schedule.GET("/teacher_fio/:teacher_fio", h.ScheduleHandler.GetScheduleByTeacherFio)
			schedule.GET("/location_name/:location_name", h.ScheduleHandler.GetScheduleByLocationName)
			scheduleAuth.PUT("/:schedule_uuid", h.ScheduleHandler.UpdateSchedule)
			scheduleAuth.DELETE("/:schedule_uuid", h.ScheduleHandler.DeleteSchedule)
		}
	}

	return routes
}
