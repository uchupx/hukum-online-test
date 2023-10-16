package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/geek-garden-test/internal/service"
	"github.com/uchupx/geek-garden-test/pkg/dto"
	"github.com/uchupx/geek-garden-test/pkg/helper"
)

type GatheringHandler struct {
	Service *service.GatheringService
}

func (h GatheringHandler) Routes(c *gin.RouterGroup) {
	c.GET("/gathering", h.GetGatherings)
	c.GET("/gathering/:id", h.GetGatheringByID)
	c.POST("/gathering", h.CreateGathering)
	c.PUT("/gathering/:id", h.UpdateGathering)
	c.DELETE("/gathering/:id", h.DeleteGathering)
}

func (h GatheringHandler) GetGatherings(c *gin.Context) {
	res, err := h.Service.GetGatherings(c)
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

func (h GatheringHandler) CreateGathering(c *gin.Context) {
	var req dto.GatheringPostV1

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	parseTime, err := helper.ParseTime(req.ScheduledAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	gathering := dto.Gathering{
		Name:        req.Name,
		Location:    req.Location,
		Creator:     req.Creator,
		ScheduledAt: *parseTime,
		Type:        req.Type,
	}

	res, err := h.Service.InsertGathering(c, gathering)

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

func (h GatheringHandler) GetGatheringByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Service.GetGatheringByID(c, helper.StringToInt64(id, 0))
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

func (h GatheringHandler) UpdateGathering(c *gin.Context) {
	var req dto.GatheringPutV1
	id := helper.StringToInt64(c.Param("id"), 0)

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	res, err := h.Service.UpdateGathering(c, id, req)
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

func (h GatheringHandler) DeleteGathering(c *gin.Context) {
	id := helper.StringToInt64(c.Param("id"), 0)

	err := h.Service.DeleteGathering(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, dto.DefaultResponse{
		Message: "success",
		Status:  http.StatusOK,
	})
}
