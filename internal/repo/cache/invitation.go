package cache

import (
	"context"

	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/internal/repo/model"
)

type InvitationCacheRepo struct {
	repo repo.InvitationRepo
}

func (r InvitationCacheRepo) GetInvitationByID(ctx context.Context, id int64) (*model.Invitation, error) {
	return r.repo.GetInvitationByID(ctx, id)
}

func (r InvitationCacheRepo) GetInvitationsByMemberID(ctx context.Context, memberID int64) ([]model.Invitation, error) {
	return r.repo.GetInvitationsByMemberID(ctx, memberID)
}

func (r InvitationCacheRepo) GetInvitationsByGatheringID(ctx context.Context, gatheringID int64) ([]model.Invitation, error) {
	return r.repo.GetInvitationsByGatheringID(ctx, gatheringID)
}

func (r InvitationCacheRepo) InsertInvitation(ctx context.Context, data model.Invitation) (*int64, error) {
	return r.repo.InsertInvitation(ctx, data)
}

func (r InvitationCacheRepo) UpdateInvitation(ctx context.Context, data model.Invitation) (*int64, error) {
	return r.repo.UpdateInvitation(ctx, data)
}

func (r InvitationCacheRepo) DeleteInvitation(ctx context.Context, data model.Invitation) (*int64, error) {
	return r.repo.DeleteInvitation(ctx, data)
}

func (r InvitationCacheRepo) GetInvitationsByPair(ctx context.Context, gatheringID, memberID int64) (*model.Invitation, error) {
	return r.repo.GetInvitationsByPair(ctx, gatheringID, memberID)
}

func NewInvitationcacheRepo(repo repo.InvitationRepo) repo.InvitationRepo {
	return &InvitationCacheRepo{
		repo: repo,
	}
}
