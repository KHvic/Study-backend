package routers

import (
	"github.com/KHvic/quiz-backend/dao"
	v1 "github.com/KHvic/quiz-backend/routers/api/v1"
	"github.com/gin-gonic/gin"
)

var router *Router

// Router ...
type Router struct {
	questionHandler *v1.QuestionHandler
}

func init() {
	router = &Router{}

	// init DAO
	questionDAO := dao.NewQuestionDAO()

	// init handlers
	questionHandler := v1.NewQuestionHandler(questionDAO)

	router.questionHandler = questionHandler
}

// InitRouter register routes
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/question/:id", router.questionHandler.GetQuestion)
	apiv1.GET("/questions/:subcat", router.questionHandler.GetSubCatQuestions)
	return r
}
