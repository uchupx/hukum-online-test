package service

import (
	"context"

	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/pkg/dto"
	errors "github.com/uchupx/geek-garden-test/pkg/error"
)

type AttendeeService struct {
	AttendeeRepo repo.AttendeeRepo
}

func (s AttendeeService) CreateAttendee(ctx context.Context, req dto.AttendeePostV1) (*int64, error) {
	isExist, err := s.AttendeeRepo.FindAttendee(ctx, req.MemberID, req.GatheringID)
	if err != nil {
		return nil, err
	}

	if isExist != nil {
		return nil, errors.ErrDataIsExist
	}

	attendee := dto.Attendee{
		MemberID:    req.MemberID,
		GatheringID: req.GatheringID,
	}

	return s.AttendeeRepo.CreateAttendee(ctx, attendee.ToModel())
}

func (s AttendeeService) CancelAttendee(ctx context.Context, req dto.AttendeePostV1) (*int64, error) {
	isExist, err := s.AttendeeRepo.FindAttendee(ctx, req.MemberID, req.GatheringID)
	if err != nil {
		return nil, err
	}

	if isExist == nil {
		return nil, errors.ErrDataNotFound
	}

	return s.AttendeeRepo.DeleteAttendee(ctx, req.MemberID, req.GatheringID)
}
