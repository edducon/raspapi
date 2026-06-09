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

// @Summary Deleting user
// @Description Deleting a user by their UUID
// @Security ApiKeyAuth
// @Tags users
// @Accept application/json
// @Produce application/json
// @Param uuid path string true "User UUID"
// @Success 200 {object} models.ResponseAPI{result=models.UserResponse} "User successfully deleted"
// @Failure 400 {object} models.ResponseAPI "Bad request format"
// @Failure 404 {object} models.ResponseAPI "User was not found"
// @Failure 409 {object} models.ResponseAPI "Conflict when deleting user"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/users/{uuid} [delete]
func (h *UsersHandler) DeleteUser(ctx *gin.Context) {
	userUUID := ctx.Param("uuid")

	deleteError := h.service.UsersService.DeleteUser(userUUID)
	if deleteError != nil {
		messageError := "Error delete user: "
		if strings.Contains(deleteError.Error(), "not found") {
			messageError += fmt.Sprintf("%s not found", userUUID)
			logger.NewErrorResponse(ctx, h.log, true, http.StatusNotFound, messageError)
		} else if strings.Contains(deleteError.Error(), "it is being referenced") {
			messageError += "target user is referenced by other records"
			logger.NewErrorResponse(ctx, h.log, true, http.StatusConflict, messageError)
		} else {
			messageError += deleteError.Error()
			logger.NewErrorResponse(ctx, h.log, true, http.StatusInternalServerError, messageError)
		}
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Delete user",
		Result: models.UserResponse{
			UserUUID: userUUID,
		},
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}
