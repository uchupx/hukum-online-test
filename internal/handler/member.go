package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/geek-garden-test/internal/service"
	"github.com/uchupx/geek-garden-test/pkg/dto"
)

type MemberHandler struct {
	Service *service.MemberService
}

func (m MemberHandler) Routes(e *gin.RouterGroup) {
	e.GET("/member", m.GetMembers)
	e.POST("/member", m.InsertMember)
	e.GET("/member/:id", m.GetMemberByID)
	e.PUT("/member/:id", m.UpdateMember)
	e.DELETE("/member/:id", m.DeleteMember)
}

func (h MemberHandler) GetMembers(c *gin.Context) {
	res, err := h.Service.GetMembers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, dto.DefaultResponse{
		Data:    res,
		Message: "success",
		Status:  http.StatusOK,
	})
}

func (h MemberHandler) InsertMember(c *gin.Context) {
	var req dto.MemberPostV1

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	member := dto.Member{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	res, err := h.Service.InsertMember(c, member)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, dto.DefaultResponse{
		Data:    res,
		Id:      *res,
		Message: "success",
		Status:  http.StatusCreated,
	})
	return
}

func (h MemberHandler) GetMemberByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid member ID",
		})
		return
	}

	// Call the GetMemberByID method on the service to get the member
	member, err := h.Service.GetMemberByID(c, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error getting member",
		})
		return
	}

	// Return a JSON response with the member and a status code of 200 (OK)
	c.JSON(http.StatusOK, dto.DefaultResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    member,
	})
	return
}

func (h MemberHandler) UpdateMember(c *gin.Context) {
	// Get the member ID from the URL parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid member ID",
		})
		return
	}

	// Bind the request body to a Member struct
	var req dto.MemberPutV1

	if err := c.Bind(&req); err != nil {
		fmt.Println("error: ", err)
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	// Call the UpdateMember method on the service to update the member
	if _, err := h.Service.UpdateMember(c, int64(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error updating member",
		})
		return
	}

	// Return a JSON response with the updated member and a status code of 200 (OK)
	c.JSON(http.StatusOK, dto.DefaultResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
	return
}

func (h MemberHandler) DeleteMember(c *gin.Context) {
	// Get the member ID from the URL parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid member ID",
		})
		return
	}

	// Call the DeleteMember method on the service to delete the member
	if err := h.Service.DeleteMember(c, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error deleting member",
		})
		return
	}

	// Return a JSON response with a status code of 200 (OK)
	c.JSON(http.StatusOK, dto.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
	})
	return
}
