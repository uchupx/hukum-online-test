package model

type Attendee struct {
	MemberID    int64 `db:"member_id"`
	GatheringID int64 `db:"gathering_id"`
}
