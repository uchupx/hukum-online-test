package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/pkg/dto"
	errors "github.com/uchupx/geek-garden-test/pkg/error"
	"github.com/uchupx/kajian-api/pkg/db"
)

type InvitationService struct {
	Repo          repo.InvitationRepo
	AttendRepo    repo.AttendeeRepo
	MemberRepo    repo.MemberRepo
	GatheringRepo repo.GatheringRepo
	DB            *db.DB
}

func (s InvitationService) CreateInvitation(ctx context.Context, data dto.Invitation) (*int64, error) {
	if err := s.checkMember(ctx, data.MemberID); err != nil {
		return nil, err
	}

	if err := s.checkGathering(ctx, data.GatheringID); err != nil {
		return nil, err
	}

	isExist, err := s.Repo.GetInvitationsByPair(ctx, data.GatheringID, data.MemberID)
	if err != nil {
		return nil, err
	}

	if isExist != nil {
		return nil, errors.ErrDataIsExist
	}

	data.Status = dto.InvitationStatusPending
	return s.Repo.InsertInvitation(ctx, data.ToModel())
}

func (s InvitationService) AttendInvitation(ctx context.Context, id int64) error {
	var inv dto.Invitation
	m, err := s.Repo.GetInvitationByID(ctx, id)
	if err != nil {
		return err
	}

	if m == nil {
		return errors.ErrDataNotFound
	}

	inv.FromModel(*m)

	// Do transaction
	err = s.DB.FTransaction(ctx, func(c context.Context, tx *sqlx.Tx) error {
		inv.Status = dto.InvitationStatusAccept
		_, err = s.Repo.UpdateInvitation(ctx, inv.ToModel())
		if err != nil {
			return err
		}

		attendee := dto.Attendee{
			GatheringID: inv.GatheringID,
			MemberID:    inv.MemberID,
		}

		_, err = s.AttendRepo.CreateAttendee(ctx, attendee.ToModel())
		if err != nil {
			return err
		}

		return nil
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s InvitationService) RejectInvitation(ctx context.Context, id int64) error {
	var inv dto.Invitation
	m, err := s.Repo.GetInvitationByID(ctx, id)
	if err != nil {
		return err
	}

	if m == nil {
		return errors.ErrDataNotFound
	}

	inv.FromModel(*m)
	inv.Status = dto.InvitationStatusReject
	_, err = s.Repo.UpdateInvitation(ctx, inv.ToModel())
	if err != nil {
		return err
	}

	return nil
}

func (h InvitationService) checkMember(ctx context.Context, memberID int64) error {
	m, err := h.MemberRepo.GetMemberByID(ctx, memberID)
	if err != nil {
		return err
	}

	if m == nil {
		return errors.ErrDataNotFound
	}

	return nil
}

func (h InvitationService) checkGathering(ctx context.Context, gatheringID int64) error {
	m, err := h.GatheringRepo.GetGatheringByID(ctx, gatheringID)
	if err != nil {
		return err
	}

	if m == nil {
		return errors.ErrDataNotFound
	}

	return nil
}

func (h InvitationService) DeleteInvitation(ctx context.Context, id int64) error {
	m, err := h.Repo.GetInvitationByID(ctx, id)
	if err != nil {
		return err
	}

	if m == nil {
		return errors.ErrDataNotFound
	}

	_, err = h.Repo.DeleteInvitation(ctx, *m)
	if err != nil {
		return err
	}

	return nil
}
