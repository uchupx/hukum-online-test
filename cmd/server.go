package cmd

import (
	"github.com/go-playground/validator/v10"
	"github.com/uchupx/geek-garden-test/internal"

	"github.com/gin-gonic/gin"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (cv *CustomValidator) SetValidator(v *validator.Validate) {
	cv.validator = v
}

func InitServer() {
	e := gin.Default()
	factoey := internal.Factory{}
	//  intialize routing
	factoey.InitRoute(e.Group("/"))
	e.Run(":8080")
}

func ServerInit() *gin.Engine {
	e := gin.Default()
	return e
}
