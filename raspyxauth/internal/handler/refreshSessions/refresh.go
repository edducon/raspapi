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

// @Summary Refreshing tokens
// @Description Refreshing tokens using a refresh token
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param request body models.RefreshRequest true "Refresh token"
// @Success 200 {object} models.ResponseAPI{result=models.AccessRefreshTokens} "Tokens successfully refreshed"
// @Failure 400 {object} models.ResponseAPI "Bad request format"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/refresh [post]
func (h *RefreshSessionsHandler) Refresh(ctx *gin.Context) {
	var request models.RefreshRequest

	messageError := "Error refresh: "
	if err := ctx.ShouldBindJSON(&request); err != nil {
		messageError += fmt.Sprintf("error parsing refresh data: %s", err.Error())
		logger.NewErrorResponse(ctx, h.log, true, http.StatusBadRequest, messageError)
		return
	}

	tokens, err := h.service.RefreshSessionsService.Refresh(&models.RefreshData{
		RefreshToken: request.RefreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		Ip:           ctx.ClientIP(),
	})
	if err != nil {
		messageError += err.Error()
		if strings.Contains(messageError, "expired") {
			logger.NewErrorResponse(ctx, h.log, true, http.StatusUnauthorized, messageError)
		} else if strings.Contains(messageError, "not found") {
			logger.NewErrorResponse(ctx, h.log, true, http.StatusUnauthorized, messageError)
		} else {
			logger.NewErrorResponse(ctx, h.log, true, http.StatusInternalServerError, messageError)
		}
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)
	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Refresh",
		Result:    tokens,
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}
