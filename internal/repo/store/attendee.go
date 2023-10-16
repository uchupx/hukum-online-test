package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/internal/repo/model"
	"github.com/uchupx/kajian-api/pkg/db"
)

const (
	attendeeInsertQuery = `INSERT INTO attendees (member_id, gathering_id) VALUES (?, ?)`
	attendeeSelectQuery = `SELECT * FROM attendees WHERE member_id = ? AND gathering_id = ?`
	attendeeDeleteQuery = `DELETE FROM attendees WHERE member_id = ? AND gathering_id = ?`
)

type AttendeeStoreRepo struct {
	db *db.DB
}

// CreateAttendee creates a new Attendee.
func (r AttendeeStoreRepo) CreateAttendee(ctx context.Context, attendee model.Attendee) (*int64, error) {
	stmt, err := r.db.FPreparexContext(ctx, attendeeInsertQuery)
	if err != nil {
		return nil, fmt.Errorf("[AttendeStoreRepo CreateAttendee] error preparing query: %v", err)
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, attendee.MemberID, attendee.GatheringID)
	if err != nil {
		return nil, fmt.Errorf("[AttendeStoreRepo CreateAttendee] error executing query: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("[AttendeStoreRepo CreateAttendee] error getting last insert id: %v", err)
	}

	return &id, nil
}

// FindAttendee finds a Attendee by member_id and gathering_id.
func (r AttendeeStoreRepo) FindAttendee(ctx context.Context, memberID, gatheringID int64) (*model.Attendee, error) {
	stmt, err := r.db.FPreparexContext(ctx, attendeeSelectQuery)
	if err != nil {
		return nil, fmt.Errorf("[AttendeStoreRepo FindAttendee] error preparing query: %v", err)
	}

	defer stmt.Close()

	var attendee model.Attendee
	if err := stmt.FQueryRowxContext(ctx, memberID, gatheringID).StructScan(&attendee); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("[AttendeStoreRepo FindAttendee] error executing query: %v", err)
	}

	return &attendee, nil
}

// DeleteAttendee deletes a Attendee by member_id and gathering_id.
func (r AttendeeStoreRepo) DeleteAttendee(ctx context.Context, memberId, gatheringId int64) (*int64, error) {
	stmt, err := r.db.FPreparexContext(ctx, attendeeDeleteQuery)
	if err != nil {
		return nil, fmt.Errorf("[AttendeStoreRepo DeleteAttendee] error preparing query: %v", err)
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, memberId, gatheringId)
	if err != nil {
		return nil, fmt.Errorf("[AttendeStoreRepo DeleteAttendee] error executing query: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("[AttendeStoreRepo DeleteAttendee] error getting last insert id: %v", err)
	}

	return &id, nil
}

func NewAttendeeStoreRepo(db *db.DB) repo.AttendeeRepo {
	return &AttendeeStoreRepo{
		db: db,
	}
}
