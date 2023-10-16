package dto

import (
	"github.com/uchupx/geek-garden-test/internal/repo/model"
	"github.com/uchupx/geek-garden-test/pkg/helper"
	"github.com/uchupx/kajian-api/pkg/logger"
)

// Member is a DTO.
type Member struct {
	ID        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     string  `json:"email"`
}

func (c *Member) FromModel(m model.Member) {

	if err := helper.MergeStruct(c, m); err != nil {
		logger.Logger.Warnf("[Member] Failed to merge struct: %v", err)
	}

}

func (c Member) ToModel() (m model.Member) {
	if err := helper.MergeStruct(&m, c); err != nil {
		logger.Logger.Warnf("[Member] Failed to merge struct: %v", err)
	}

	return m
}
