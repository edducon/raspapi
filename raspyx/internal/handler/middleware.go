package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"raspyx2/internal/handler/constHandler"
	"raspyx2/internal/models"
	"raspyx2/pkg/logger"
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

		status, body, err := sendHTTPRequest(httpRequestParams{
			Method: http.MethodGet,
			URL:    fmt.Sprintf("http://%s/%s", h.cfg.Auth.BasePath, h.cfg.Auth.VerifyPath),
			Headers: map[string]string{
				"Authorization": authHeader,
			},
		})
		if err != nil {
			logger.NewErrorResponse(c, h.log, true, http.StatusUnauthorized, messageError+err.Error())
			return
		}

		var res models.ResponseAPI
		err = json.Unmarshal(body, &res)
		if err != nil {
			logger.NewErrorResponse(c, h.log, true, http.StatusUnauthorized, messageError+err.Error())
			return
		}

		if status != 200 {
			logger.NewErrorResponse(c, h.log, true, http.StatusUnauthorized, res.Errors.Message)
			return
		}

		mRes, _ := res.Result.(map[string]interface{})
		c.Set("userId", mRes["sub"])
		c.Set("access_level", mRes["access_level"])

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

type httpRequestParams struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
	Timeout time.Duration
}

func sendHTTPRequest(params httpRequestParams) (int, []byte, error) {
	if params.Method == "" {
		params.Method = http.MethodGet
	}
	if params.Timeout == 0 {
		params.Timeout = 10 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), params.Timeout)
	defer cancel()

	var bodyReader io.Reader
	if len(params.Body) > 0 {
		bodyReader = bytes.NewReader(params.Body)
	}

	req, err := http.NewRequestWithContext(ctx, params.Method, params.URL, bodyReader)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request: %w", err)
	}

	for key, value := range params.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: params.Timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return resp.StatusCode, body, nil
}
