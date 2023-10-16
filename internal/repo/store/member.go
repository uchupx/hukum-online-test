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
	// GetMemberByIDQuery is a query to get member by id
	GetMemberByIDQuery = `SELECT id, first_name, last_name, email FROM members WHERE id = ?`

	// GetMembersQuery is a query to get all members
	GetMembersQuery = `SELECT id, first_name, last_name, email FROM members`

	// InsertMemberQuery is a query to insert member
	InsertMemberQuery = `INSERT INTO members (first_name, last_name, email) VALUES (?, ?, ?)`

	// UpdateMemberQuery is a query to update member
	UpdateMemberQuery = `UPDATE members SET first_name = ?, last_name = ?, email = ? WHERE id = ?`

	// DeleteMemberQuery is a query to delete member
	DeleteMemberQuery = `DELETE FROM members WHERE id = ?`
)

type MemberStoreRepo struct {
	db *db.DB
}

func (c MemberStoreRepo) GetMemberByID(ctx context.Context, id int64) (*model.Member, error) {
	var item model.Member

	stmt, err := c.db.FPreparexContext(ctx, GetMemberByIDQuery)
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo GetMemberByID] error prepare statement: %w", err)
	}

	defer stmt.Close()

	if err = stmt.FQueryRowxContext(ctx, id).StructScan(&item); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo GetMemberByID] error scan: %w", err)
	}

	return &item, nil
}

func (c MemberStoreRepo) GetMembers(ctx context.Context) ([]model.Member, error) {
	var items []model.Member

	stmt, err := c.db.FPreparexContext(ctx, GetMembersQuery)
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo GetMembers] error prepare statement: %w", err)
	}

	defer stmt.Close()

	rows, err := stmt.FQueryxContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo GetMembers] error query: %w", err)
	}

	for rows.Next() {
		var item model.Member

		if err = rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("[MemberStoreRepo GetMembers] error scan: %w", err)
		}

		items = append(items, item)
	}

	return items, nil
}

func (c MemberStoreRepo) InsertMember(ctx context.Context, data model.Member) (*int64, error) {
	stmt, err := c.db.FPreparexContext(ctx, InsertMemberQuery)
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo InsertMember] error prepare statement: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, data.FirstName, data.LastName, data.Email)
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo InsertMember] error exec: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo InsertMember] error get last insert id: %w", err)
	}

	return &id, nil
}

func (c MemberStoreRepo) UpdateMember(ctx context.Context, data model.Member) (*int64, error) {
	stmt, err := c.db.FPreparexContext(ctx, UpdateMemberQuery)
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo UpdateMember] error prepare statement: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, data.FirstName, data.LastName, data.Email, data.ID)
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo UpdateMember] error exec: %w", err)
	}

	rowAfftected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo UpdateMember] error get rows affected: %w", err)
	}

	return &rowAfftected, nil
}

func (c MemberStoreRepo) DeleteMember(ctx context.Context, data model.Member) (*int64, error) {
	stmt, err := c.db.FPreparexContext(ctx, DeleteMemberQuery)
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo DeleteMember] error prepare statement: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, data.ID)
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo DeleteMember] error exec: %w", err)
	}

	rowAfftected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("[MemberStoreRepo DeleteMember] error get rows affected: %w", err)
	}

	return &rowAfftected, nil
}

func NewMemberStoreRepo(db *db.DB) repo.MemberRepo {
	return &MemberStoreRepo{db: db}
}
