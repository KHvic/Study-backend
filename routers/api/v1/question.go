package v1

import (
	"net/http"

	"github.com/KHvic/quiz-backend/dao"
	"github.com/KHvic/quiz-backend/pkg/app"
	"github.com/KHvic/quiz-backend/pkg/constant"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// QuestionHandler ...
type QuestionHandler struct {
	questionDAO dao.QuestionDAO
}

// NewQuestionHandler return question handler
func NewQuestionHandler(questionDAO dao.QuestionDAO) *QuestionHandler {
	return &QuestionHandler{questionDAO: questionDAO}
}

// GetQuestion ...
func (h *QuestionHandler) GetQuestion(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt64()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, constant.BadRequest, nil)
		return
	}

	question, err := h.questionDAO.GetByID(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.InternalError, nil)
		return
	}
	appG.Response(http.StatusOK, constant.Success, question)
}

// GetSubCatQuestions ...
func (h *QuestionHandler) GetSubCatQuestions(c *gin.Context) {
	appG := app.Gin{C: c}
	count := com.StrTo(c.DefaultQuery("count", "1")).MustInt()
	subcat := com.StrTo(c.DefaultQuery("subcat", "")).String()
	valid := validation.Validation{}
	valid.Min(count, 1, "count")
	valid.Max(count, 20, "count")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, constant.BadRequest, nil)
		return
	}

	questions, err := h.questionDAO.MGetBySubCat(subcat, count)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.InternalError, nil)
		return
	}
	appG.Response(http.StatusOK, constant.Success, questions)
}

// Validate ...
func (h *QuestionHandler) Validate(c *gin.Context) {
	appG := app.Gin{C: c}
	start := com.StrTo(c.DefaultQuery("start", "0")).MustInt64()
	count := com.StrTo(c.DefaultQuery("count", "10")).MustInt64()
	valid := validation.Validation{}
	valid.Min(count, 1, "count")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, constant.BadRequest, nil)
		return
	}

	questions, err := h.questionDAO.GetByOffsetAndLimit(start, count)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.InternalError, nil)
		return
	}
	var failedQuestionIDs []int64
	for _, q := range questions {
		if q.Validate() != nil {
			failedQuestionIDs = append(failedQuestionIDs, q.ID)
		}
	}

	appG.Response(http.StatusOK, constant.Success, failedQuestionIDs)
}
