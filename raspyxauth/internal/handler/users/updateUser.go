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
	"strings"
)

// @Summary Updating user
// @Description Updating user information by their UUID
// @Security ApiKeyAuth
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param uuid path string true "User uuid"
// @Param request body models.UpdateUserRequest true "User update data"
// @Success 200 {object} models.ResponseAPI{result=models.UserResponse} "User data has been successfully updated"
// @Failure 400 {object} models.ResponseAPI "Bad request format"
// @Failure 404 {object} models.ResponseAPI "User was not found"
// @Failure 409 {object} models.ResponseAPI "Conflict when updating user"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/users/{uuid} [put]
func (h *UsersHandler) UpdateUser(ctx *gin.Context) {
	var request models.UpdateUserRequest
	userUUID := ctx.Param("uuid")

	messageError := "Error update user: "
	if err := ctx.ShouldBindJSON(&request); err != nil {
		messageError += fmt.Sprintf("error parsing user data: %s", err.Error())
		logger.NewErrorResponse(ctx, h.log, true, http.StatusBadRequest, messageError)
		return
	}

	updateError := h.service.UsersService.UpdateUser(userUUID, &request)
	if updateError != nil {
		if errors.Is(updateError, sql.ErrNoRows) {
			messageError += fmt.Sprintf("%s not found", userUUID)
			logger.NewErrorResponse(ctx, h.log, true, http.StatusNotFound, messageError)
		} else if strings.Contains(updateError.Error(), "already exists") {
			messageError += "target user already exists"
			logger.NewErrorResponse(ctx, h.log, true, http.StatusConflict, messageError)
		} else {
			messageError += updateError.Error()
			logger.NewErrorResponse(ctx, h.log, true, http.StatusInternalServerError, messageError)
		}
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Update user",
		Result: models.UserResponse{
			UserUUID: userUUID,
		},
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}
