package users

import (
	"JWTAuth/internal/handler/constHandler"
	"JWTAuth/internal/models"
	"JWTAuth/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @Summary Creating a user
// @Description Creating a new user
// @Security ApiKeyAuth
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param request body models.AddUserRequest true "User data"
// @Success 201 {object} models.ResponseAPI{result=models.UserResponse} "User successfully created"
// @Failure 400 {object} models.ResponseAPI "Bad request format"
// @Failure 409 {object} models.ResponseAPI "Conflict when creating user"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/users [post]
func (h *UsersHandler) CreateUser(ctx *gin.Context) {
	var request models.AddUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		messageError := fmt.Sprintf("Error parsing user data: %s", err.Error())
		logger.NewErrorResponse(ctx, h.log, true, http.StatusBadRequest, messageError)
		return
	}

	request.Username = strings.TrimSpace(request.Username)

	createUserUUID, createUserError := h.service.UsersService.CreateUser(&request)
	if createUserError != nil {
		messageError := fmt.Sprintf("Error new user: %s", createUserError.Error())
		if strings.Contains(createUserError.Error(), "already exists") {
			logger.NewErrorResponse(ctx, h.log, true, http.StatusConflict, messageError)
		} else if strings.Contains(createUserError.Error(), "validate") {
			logger.NewErrorResponse(ctx, h.log, true, http.StatusBadRequest, messageError)
		} else {
			logger.NewErrorResponse(ctx, h.log, true, http.StatusInternalServerError, messageError)
		}
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)
	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Add new user",
		Result: models.UserResponse{
			UserUUID: createUserUUID,
		},
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusCreated, responseApi)
}
