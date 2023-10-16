package dto

type AttendeePostV1 struct {
	MemberID    int64 `json:"member_id" binding:"required"`
	GatheringID int64 `json:"gathering_id" binding:"required"`
}
