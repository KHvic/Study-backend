package v1

import (
	"net/http"

	"github.com/KHvic/study-backend/dao"
	"github.com/KHvic/study-backend/pkg/app"
	"github.com/KHvic/study-backend/pkg/constant"
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
	k := com.StrTo(c.DefaultQuery("k", "1")).MustInt()
	subcat := com.StrTo(c.Param("subcat")).String()
	valid := validation.Validation{}
	valid.Min(k, 1, "k")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, constant.BadRequest, nil)
		return
	}

	questions, err := h.questionDAO.GetBySubCatRandK(subcat, k)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.InternalError, nil)
		return
	}
	appG.Response(http.StatusOK, constant.Success, questions)
}
