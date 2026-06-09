package refreshSessions

import (
	"JWTAuth/internal/handler/constHandler"
	"JWTAuth/internal/models"
	"JWTAuth/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Logging in
// @Description Getting JWT token using credentials
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param request body models.LoginData true "Login data"
// @Success 200 {object} models.ResponseAPI{result=models.AccessRefreshTokens} "RefreshSession successfully created"
// @Failure 400 {object} models.ResponseAPI "Bad request format"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/login [post]
func (h *RefreshSessionsHandler) Login(ctx *gin.Context) {
	var request models.LoginData

	messageError := "Error login: "
	if err := ctx.ShouldBindJSON(&request); err != nil {
		messageError += fmt.Sprintf("error parsing login data: %s", err.Error())
		logger.NewErrorResponse(ctx, h.log, true, http.StatusBadRequest, messageError)
		return
	}

	user, err := h.service.UsersService.Login(&request)
	if err != nil {
		messageError += err.Error()
		logger.NewErrorResponse(ctx, h.log, true, http.StatusUnauthorized, messageError)
		return
	}

	tokens, err := h.service.RefreshSessionsService.CreateJWT(&models.CreateJWT{
		UserUUID:    user.UUID,
		UserAgent:   ctx.Request.UserAgent(),
		Ip:          ctx.ClientIP(),
		AccessLevel: user.AccessLevel,
		Config:      h.cfg,
	})
	if err != nil {
		messageError += err.Error()
		logger.NewErrorResponse(ctx, h.log, true, http.StatusUnauthorized, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)
	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Login",
		Result:    tokens,
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}
