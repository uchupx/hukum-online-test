package dto

type InvitationPostV1 struct {
	MemberID    int64 `json:"member_id" binding:"required"`
	GatheringID int64 `json:"gathering_id" binding:"required"`
}

type InvitationPutV1 struct {
	Status string `json:"status" binding:"required"`
}
