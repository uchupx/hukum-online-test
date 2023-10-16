package cache

import (
	"context"

	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/internal/repo/model"
)

type MemberCacheRepo struct {
	repo repo.MemberRepo
}

func (c MemberCacheRepo) GetMemberByID(ctx context.Context, id int64) (*model.Member, error) {
	return c.repo.GetMemberByID(ctx, id)
}

func (c MemberCacheRepo) GetMembers(ctx context.Context) ([]model.Member, error) {
	return c.repo.GetMembers(ctx)
}

func (c MemberCacheRepo) InsertMember(ctx context.Context, data model.Member) (*int64, error) {
	return c.repo.InsertMember(ctx, data)
}

func (c MemberCacheRepo) UpdateMember(ctx context.Context, data model.Member) (*int64, error) {
	return c.repo.UpdateMember(ctx, data)
}

func (c MemberCacheRepo) DeleteMember(ctx context.Context, data model.Member) (*int64, error) {
	return c.repo.DeleteMember(ctx, data)
}

func NewMemberCacheRepo(repo repo.MemberRepo) *MemberCacheRepo {
	return &MemberCacheRepo{
		repo: repo,
	}
}
