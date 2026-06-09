package refreshSessions

import (
	"JWTAuth/internal/handler/constHandler"
	"JWTAuth/internal/models"
	"JWTAuth/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @Summary User registration
// @Description Adding user to db
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param request body models.RegData true "Registration data"
// @Success 200 {object} models.ResponseAPI{result=models.UserResponse} "User successfully created"
// @Failure 400 {object} models.ResponseAPI "Bad request format"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/reg [post]
func (h *RefreshSessionsHandler) Registration(ctx *gin.Context) {
	var request models.RegData

	messageError := "Error registration: "
	if err := ctx.ShouldBindJSON(&request); err != nil {
		messageError += fmt.Sprintf("error parsing registration data: %s", err.Error())
		logger.NewErrorResponse(ctx, h.log, true, http.StatusBadRequest, messageError)
		return
	}

	userUUID, err := h.service.UsersService.CreateUser(&models.AddUserRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		messageError += err.Error()
		if strings.Contains(err.Error(), "validate") {
			logger.NewErrorResponse(ctx, h.log, true, http.StatusBadRequest, messageError)
		} else if strings.Contains(err.Error(), "already exists") {
			logger.NewErrorResponse(ctx, h.log, true, http.StatusConflict, messageError)
		} else {
			logger.NewErrorResponse(ctx, h.log, true, http.StatusUnauthorized, messageError)
		}
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)
	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Successful registration",
		Result:    &models.UserResponse{UserUUID: userUUID},
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}
