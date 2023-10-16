package cache

import (
	"context"

	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/internal/repo/model"
)

type GatheringCacheRepo struct {
	repo repo.GatheringRepo
}

func (r GatheringCacheRepo) GetGatheringByID(ctx context.Context, id int64) (*model.Gathering, error) {
	return r.repo.GetGatheringByID(ctx, id)
}

func (r GatheringCacheRepo) GetGatherings(ctx context.Context) ([]model.Gathering, error) {
	return r.repo.GetGatherings(ctx)
}

func (r GatheringCacheRepo) InsertGathering(ctx context.Context, data model.Gathering) (*int64, error) {
	return r.repo.InsertGathering(ctx, data)
}

func (r GatheringCacheRepo) UpdateGathering(ctx context.Context, data model.Gathering) (*int64, error) {
	return r.repo.UpdateGathering(ctx, data)
}

func (r GatheringCacheRepo) DeleteGathering(ctx context.Context, data model.Gathering) (*int64, error) {
	return r.repo.DeleteGathering(ctx, data)
}

func NewGatheringcacheRepo(repo repo.GatheringRepo) repo.GatheringRepo {
	return &GatheringCacheRepo{
		repo: repo,
	}
}
