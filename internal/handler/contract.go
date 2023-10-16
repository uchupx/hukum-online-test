package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Routes(e *gin.RouterGroup)
}
