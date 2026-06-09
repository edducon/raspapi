package users

import (
	"JWTAuth/internal/handler/constHandler"
	"JWTAuth/internal/models"
	"JWTAuth/pkg/logger"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Getting users
// @Description Getting a list of all users
// @Security ApiKeyAuth
// @Tags users
// @Accept application/json
// @Produce application/json
// @Success 200 {object} models.ResponseAPI{result=[]models.User} "User data has been successfully received"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/users [get]
func (h *UsersHandler) GetUsers(ctx *gin.Context) {
	GetAllUsers, GetAllUsersError := h.service.UsersService.GetUsers()

	if GetAllUsersError != nil {
		messageError := fmt.Sprintf("Error get users: %s", GetAllUsersError.Error())
		logger.NewErrorResponse(ctx, h.log, true, http.StatusInternalServerError, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)
	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Get users",
		Result:    GetAllUsers,
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}

// @Summary Getting a user by UUID
// @Description Getting information about a user by their UUID
// @Security ApiKeyAuth
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param uuid path string true "User UUID"
// @Success 200 {object} models.ResponseAPI{result=models.User} "User data has been successfully received"
// @Failure 404 {object} models.ResponseAPI "User was not found"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/users/uuid/{uuid} [get]
func (h *UsersHandler) GetUserByUUID(ctx *gin.Context) {
	userUUID := ctx.Param("uuid")
	getUser, getUserError := h.service.UsersService.GetUserByUUID(userUUID)
	if getUserError != nil {
		messageError := "Error get user: "
		if errors.Is(getUserError, sql.ErrNoRows) {
			messageError += fmt.Sprintf("user not found: %s", userUUID)
			logger.NewErrorResponse(ctx, h.log, true, http.StatusNotFound, messageError)
		} else {
			messageError += fmt.Sprintf("%s: %s", userUUID, getUserError.Error())
			logger.NewErrorResponse(ctx, h.log, true, http.StatusInternalServerError, messageError)
		}
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Get user",
		Result:    getUser,
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}

// @Summary Getting a user by username
// @Description Getting information about a user by their username
// @Security ApiKeyAuth
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param username path string true "username"
// @Success 200 {object} models.ResponseAPI{result=[]models.User} "User data has been successfully received"
// @Failure 404 {object} models.ResponseAPI "User was not found"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/users/username/{username} [get]
func (h *UsersHandler) GetUsersByUsername(ctx *gin.Context) {
	username := strings.TrimSpace(ctx.Param("username"))
	getUser, getUserError := h.service.UsersService.GetUsersByUsername(username)
	if getUserError != nil {
		messageError := "Error get user: "
		if errors.Is(getUserError, sql.ErrNoRows) {
			messageError += fmt.Sprintf("user not found: %s", username)
			logger.NewErrorResponse(ctx, h.log, true, http.StatusNotFound, messageError)
		} else {
			messageError += fmt.Sprintf("%s: %s", username, getUserError.Error())
			logger.NewErrorResponse(ctx, h.log, true, http.StatusInternalServerError, messageError)
		}
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Get user",
		Result:    getUser,
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}

// @Summary Получение пользователей по уровню доступа
// @Description Getting information about users whose access level is lower or equal to the specified one
// @Security ApiKeyAuth
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param access_level path string true "User access level"
// @Success 200 {object} models.ResponseAPI{result=[]models.User} "User data has been successfully received"
// @Failure 404 {object} models.ResponseAPI "User was not found"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/users/accesslevel/{access_level} [get]
func (h *UsersHandler) GetUsersByAccessLevel(ctx *gin.Context) {
	al := strings.TrimSpace(ctx.Param("access_level"))
	accessLevel, errParseInt := strconv.ParseInt(al, 10, 64)
	messageError := "Error get user: "
	if errParseInt != nil {
		messageError += errParseInt.Error()
		logger.NewErrorResponse(ctx, h.log, true, http.StatusBadRequest, messageError)
		return
	}

	if int(ctx.GetFloat64("access_level")) < int(accessLevel) {
		messageError += "you cannot get users with an access level higher than yours"
		logger.NewErrorResponse(ctx, h.log, true, http.StatusForbidden, messageError)
		return
	}

	getUser, getUserError := h.service.UsersService.GetUsersByAccessLevel(int(accessLevel))
	if getUserError != nil {
		if errors.Is(getUserError, sql.ErrNoRows) {
			messageError += fmt.Sprintf("user not found: access_level=%s", al)
			logger.NewErrorResponse(ctx, h.log, true, http.StatusNotFound, messageError)
		} else {
			messageError += fmt.Sprintf("access_level=%s: %s", al, getUserError.Error())
			logger.NewErrorResponse(ctx, h.log, true, http.StatusInternalServerError, messageError)
		}
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Get user",
		Result:    getUser,
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}
