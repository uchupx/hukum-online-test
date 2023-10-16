package dto

import (
	"github.com/uchupx/geek-garden-test/internal/repo/model"
	"github.com/uchupx/geek-garden-test/pkg/helper"
	"github.com/uchupx/kajian-api/pkg/logger"
)

const (
	InvitationStatusPending = "PENDING"
	InvitationStatusAccept  = "ACCEPT"
	InvitationStatusReject  = "REJECT"
)

type Invitation struct {
	ID          int64  `json:"id"`
	MemberID    int64  `json:"member_id"`
	GatheringID int64  `json:"gathering_id"`
	Status      string `json:"status"`
}

func (c *Invitation) FromModel(m model.Invitation) {
	if err := helper.MergeStruct(c, m); err != nil {
		logger.Logger.Warnf("[Invitation] Failed to merge struct: %v", err)
	}

}

func (c Invitation) ToModel() (m model.Invitation) {
	if err := helper.MergeStruct(&m, c); err != nil {
		logger.Logger.Warnf("[Invitation] Failed to merge struct: %v", err)
	}

	return m
}
