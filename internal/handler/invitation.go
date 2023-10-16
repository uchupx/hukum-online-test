package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/geek-garden-test/internal/service"
	"github.com/uchupx/geek-garden-test/pkg/dto"
	"github.com/uchupx/geek-garden-test/pkg/helper"
)

type InvitationHandler struct {
	Service *service.InvitationService
}

func (h InvitationHandler) Routes(router *gin.RouterGroup) {
	// router.GET("/invitation", h.GetInvitations)
	router.POST("/invitation", h.CreateInvitation)
	router.PUT("/invitation/:id/accept", h.AcceptInvitation)
	router.PUT("/invitation/:id/reject", h.RejectInvitation)
	router.DELETE("/invitation/:id", h.DeleteInvitation)
}

func (h InvitationHandler) CreateInvitation(c *gin.Context) {
	var data dto.Invitation
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	id, err := h.Service.CreateInvitation(c.Request.Context(), data)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.DefaultResponse{
		Data:    id,
		Message: "success",
		Status:  http.StatusCreated,
		Id:      *id,
	})
}

func (h InvitationHandler) AcceptInvitation(c *gin.Context) {
	id := helper.StringToInt64(c.Param("id"), 0)

	err := h.Service.AttendInvitation(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, dto.DefaultResponse{
		Message: "success",
		Status:  http.StatusOK,
	})
}

func (h InvitationHandler) RejectInvitation(c *gin.Context) {
	id := helper.StringToInt64(c.Param("id"), 0)

	err := h.Service.RejectInvitation(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, dto.DefaultResponse{
		Message: "success",
		Status:  http.StatusOK,
	})
}

func (h InvitationHandler) DeleteInvitation(c *gin.Context) {
	id := helper.StringToInt64(c.Param("id"), 0)

	err := h.Service.DeleteInvitation(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, dto.DefaultResponse{
		Message: "success",
		Status:  http.StatusOK,
	})
}
