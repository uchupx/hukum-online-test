package repo

import (
	"context"

	"github.com/uchupx/geek-garden-test/internal/repo/model"
)

// Attendee is the interface that wraps the basic Attendee methods.
type AttendeeRepo interface {
	// CreateAttendee creates a new Attendee.
	CreateAttendee(ctx context.Context, attendee model.Attendee) (*int64, error)
	FindAttendee(ctx context.Context, memberID, gatheringID int64) (*model.Attendee, error)
	DeleteAttendee(ctx context.Context, memberId, gatheringId int64) (*int64, error)
}
