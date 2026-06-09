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

// @Summary Token verification
// @Description Verification of the token in the Authorization header
// @Security ApiKeyAuth
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} models.ResponseAPI{result=map[string]any} "Token is valid"
// @Failure 400 {object} models.ResponseAPI "Bad request format"
// @Failure 401 {object} models.ResponseAPI "Token is invalid"
// @Failure 500 {object} models.ResponseAPI "Internal server error"
// @Router /v1/verify [get]
func (h *RefreshSessionsHandler) Verify(ctx *gin.Context) {
	messageError := "Error verify: "

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		messageError += "authorization header required"
		logger.NewErrorResponse(ctx, h.log, true, http.StatusUnauthorized, messageError)
		return
	}

	claims, err := h.service.RefreshSessionsService.VerifyJWT(authHeader)
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
		Result:    claims,
	}

	h.log.Debug(fmt.Sprintf("response %s: %+v", ctx.Request.URL.Path, responseApi))

	ctx.JSON(http.StatusOK, responseApi)
}
