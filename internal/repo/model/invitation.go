package model

type Invitation struct {
	ID          int64  `db:"id"`
	MemberID    int64  `db:"member_id"`
	GatheringID int64  `db:"gathering_id"`
	Status      string `db:"status"`
}
