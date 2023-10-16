package cache

import (
	"context"

	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/internal/repo/model"
)

type AttendeeCacheRepo struct {
	repo repo.AttendeeRepo
}

// CreateAttendee creates a new Attendee.
func (r AttendeeCacheRepo) CreateAttendee(ctx context.Context, attendee model.Attendee) (*int64, error) {
	return r.repo.CreateAttendee(ctx, attendee)
}

// FindAttendee finds a Attendee by member_id and gathering_id.
func (r AttendeeCacheRepo) FindAttendee(ctx context.Context, memberID, gatheringID int64) (*model.Attendee, error) {
	return r.repo.FindAttendee(ctx, memberID, gatheringID)
}

// DeleteAttendee deletes a Attendee.
func (r AttendeeCacheRepo) DeleteAttendee(ctx context.Context, memberId, gatheringId int64) (*int64, error) {
	return r.repo.DeleteAttendee(ctx, memberId, gatheringId)
}

func NewAttendeecacheRepo(repo repo.AttendeeRepo) repo.AttendeeRepo {
	return &AttendeeCacheRepo{
		repo: repo,
	}
}
