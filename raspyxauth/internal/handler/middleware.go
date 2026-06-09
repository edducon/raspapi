package handler

import (
	"JWTAuth/internal/handler/constHandler"
	"JWTAuth/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		xRequestID := c.GetHeader("X-Request-ID")

		if xRequestID == "" {
			xRequestID = uuid.NewString()
		}

		c.Set(constHandler.REQUEST_ID, xRequestID)

		t := time.Now()

		c.Next()

		latency := time.Since(t)

		currRequestId, _ := c.Get(constHandler.REQUEST_ID)

		h.log.Info(fmt.Sprintf("%d %s %s %s %s %s",
			c.Writer.Status(),
			currRequestId,
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			latency),
		)
	}
}

func (h *Handler) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		messageError := "Error verify: "

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			messageError += "authorization header required"
			logger.NewErrorResponse(c, h.log, true, http.StatusUnauthorized, messageError)
			return
		}

		claims, err := h.services.RefreshSessionsService.VerifyJWT(authHeader)
		if err != nil {
			messageError += err.Error()
			logger.NewErrorResponse(c, h.log, true, http.StatusUnauthorized, messageError)
			return
		}

		if claims == nil {
			messageError += "empty claims"
			logger.NewErrorResponse(c, h.log, true, http.StatusUnauthorized, messageError)
			return
		}

		c.Set("username", (*claims)["sub"])
		c.Set("access_level", (*claims)["access_level"])

		c.Next()
	}
}

func (h *Handler) CheckAL(accessLevel int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAccessLevel := int(c.GetFloat64("access_level"))
		if userAccessLevel < accessLevel {
			logger.NewErrorResponse(c, h.log, true, http.StatusForbidden, "forbidden")
			return
		}

		c.Next()
	}
}
