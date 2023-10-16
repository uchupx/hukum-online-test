package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/geek-garden-test/internal/service"
	"github.com/uchupx/geek-garden-test/pkg/dto"
)

type AttendeeHandler struct {
	Service *service.AttendeeService
}

func (h AttendeeHandler) Routes(e *gin.RouterGroup) {
	e.POST("/attendee", h.CreateAttendee)
	e.POST("/attendee/cancel", h.CancelAttende)
}

func (h AttendeeHandler) CreateAttendee(c *gin.Context) {
	var attendee dto.AttendeePostV1
	if err := c.ShouldBindJSON(&attendee); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	res, err := h.Service.CreateAttendee(c, attendee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, dto.DefaultResponse{
		Data:    res,
		Message: "success",
		Status:  http.StatusCreated,
		Id:      *res,
	})
}

func (h AttendeeHandler) CancelAttende(c *gin.Context) {
	var attendee dto.AttendeePostV1
	if err := c.ShouldBindJSON(&attendee); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	res, err := h.Service.CancelAttendee(c, attendee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, dto.DefaultResponse{
		Data:    res,
		Message: "success",
		Status:  http.StatusOK,
		Id:      *res,
	})
}
