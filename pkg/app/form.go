package app

import (
	"net/http"

	"github.com/KHvic/study-backend/pkg/constant"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, constant.BadRequest
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, constant.InternalError
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, constant.BadRequest
	}

	return http.StatusOK, constant.Success
}
