package repo

import (
	"context"

	"github.com/uchupx/geek-garden-test/internal/repo/model"
)

type MemberRepo interface {
	GetMemberByID(ctx context.Context, id int64) (*model.Member, error)
	GetMembers(ctx context.Context) ([]model.Member, error)
	InsertMember(ctx context.Context, data model.Member) (*int64, error)
	UpdateMember(ctx context.Context, data model.Member) (*int64, error)
	DeleteMember(ctx context.Context, data model.Member) (*int64, error)
}
