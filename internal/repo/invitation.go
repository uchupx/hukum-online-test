package repo

import (
	"context"

	"github.com/uchupx/geek-garden-test/internal/repo/model"
)

type InvitationRepo interface {
	GetInvitationByID(ctx context.Context, id int64) (*model.Invitation, error)
	GetInvitationsByMemberID(ctx context.Context, memberID int64) ([]model.Invitation, error)
	GetInvitationsByGatheringID(ctx context.Context, gatheringID int64) ([]model.Invitation, error)
	GetInvitationsByPair(ctx context.Context, gatheringID, memberID int64) (*model.Invitation, error)
	InsertInvitation(ctx context.Context, data model.Invitation) (*int64, error)
	UpdateInvitation(ctx context.Context, data model.Invitation) (*int64, error)
	DeleteInvitation(ctx context.Context, data model.Invitation) (*int64, error)
}
