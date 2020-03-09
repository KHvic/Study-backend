package v1

import (
	"net/http"

	"github.com/KHvic/quiz-backend/pkg/app"
	"github.com/KHvic/quiz-backend/pkg/constant"
	"github.com/gin-gonic/gin"
)

// QuestionHandler ...
type HealthCheckHandler struct {}

// NewQuestionHandler return question handler
func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// GetQuestion ...
func (h *HealthCheckHandler) Get(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, constant.Success, "healthy!")
}
