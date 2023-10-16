package dto

import (
	"time"

	"github.com/uchupx/geek-garden-test/internal/repo/model"
	"github.com/uchupx/geek-garden-test/pkg/helper"
	"github.com/uchupx/kajian-api/pkg/logger"
)

type Gathering struct {
	ID          int64     `json:"id"`
	Creator     int64     `json:"creator"`
	Location    string    `json:"location"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
}

func (c *Gathering) FromModel(m model.Gathering) {
	if err := helper.MergeStruct(c, m); err != nil {
		logger.Logger.Warnf("[Gathering] Failed to merge struct: %v", err)
	}

}

func (c Gathering) ToModel() (m model.Gathering) {
	if err := helper.MergeStruct(&m, c); err != nil {
		logger.Logger.Warnf("[Gathering] Failed to merge struct: %v", err)
	}

	return m
}
