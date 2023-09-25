package v1

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kafka_example/api-gateway/api/handler/v1/http"
	"github.com/kafka_example/api-gateway/api/tokens"
	"github.com/kafka_example/api-gateway/config"
	"github.com/kafka_example/api-gateway/pkg/logger"
	"github.com/kafka_example/api-gateway/services"
	"github.com/stretchr/testify/assert"
)

type handlerV1 struct {
	assert.Assertions
	Cfg            *config.Config
	Log            logger.Logger
	JwtHandler     tokens.JWTHandler
	serviceManager services.ServiceManagerI
	Testing        testing.T
}

type HandlerV1Option struct {
	Cfg            *config.Config
	Log            logger.Logger
	JwtHandler     tokens.JWTHandler
	ServiceManager services.ServiceManagerI
	Testing        testing.T
}

func New(optoins *HandlerV1Option) *handlerV1 {
	return &handlerV1{
		Cfg:            optoins.Cfg,
		Log:            optoins.Log,
		JwtHandler:     optoins.JwtHandler,
		serviceManager: optoins.ServiceManager,
		Testing:        optoins.Testing,
	}
}

func (h *handlerV1) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.Log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.Log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.Log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}
