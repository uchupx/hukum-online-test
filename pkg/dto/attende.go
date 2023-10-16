package dto

import (
	"github.com/uchupx/geek-garden-test/internal/repo/model"
	"github.com/uchupx/geek-garden-test/pkg/helper"
	"github.com/uchupx/kajian-api/pkg/logger"
)

type Attendee struct {
	MemberID    int64 `json:"member_id"`
	GatheringID int64 `json:"gathering_id"`
}

func (a Attendee) ToModel() (m model.Attendee) {
	if err := helper.MergeStruct(&m, a); err != nil {
		logger.Logger.Warnf("[Attendee] Failed to merge struct: %v", err)
	}

	return m
}
