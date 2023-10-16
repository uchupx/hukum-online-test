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
	invitationGetByIDQuery          = "SELECT id, member_id, gathering_id, status FROM invitations WHERE id = ?"
	invitationGetByMemberIDQuery    = "SELECT id, member_id, gathering_id, status FROM invitations WHERE member_id = ?"
	invitationGetByGatheringIDQuery = "SELECT id, member_id, gathering_id, status FROM invitations WHERE gathering_id = ?"
	invitationGetByPair             = "SELECT id, member_id, gathering_id, status FROM invitations WHERE gathering_id = ? AND member_id = ?"
	invitationInsertQuery           = "INSERT INTO invitations (member_id, gathering_id, status) VALUES (?, ?, ?)"
	invitationUpdateQuery           = "UPDATE invitations SET member_id = ?, gathering_id = ?, status = ? WHERE id = ?"
	invitationDeleteQuery           = "DELETE FROM invitations WHERE id = ?"
)

type InvitationStoreRepo struct {
	db *db.DB
}

func (r InvitationStoreRepo) GetInvitationByID(ctx context.Context, id int64) (*model.Invitation, error) {
	stmt, err := r.db.FPreparexContext(ctx, invitationGetByIDQuery)
	if err != nil {
		fmt.Println("[InvitationRepo GetInvitationByID] error: ", err)
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmt.Close()

	var invitation model.Invitation
	if err := stmt.FQueryRowxContext(ctx, id).StructScan(&invitation); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		fmt.Println("[InvitationRepo GetInvitationByID] error: ", err)
		return nil, err
	}

	return &invitation, nil
}

func (r InvitationStoreRepo) GetInvitationsByMemberID(ctx context.Context, memberID int64) ([]model.Invitation, error) {
	stmt, err := r.db.FPreparexContext(ctx, invitationGetByMemberIDQuery)
	if err != nil {
		fmt.Println("[InvitationRepo GetInvitationsByMemberID] error: ", err)
		return nil, err
	}

	defer stmt.Close()

	var invitations []model.Invitation
	rows, err := stmt.FQueryxContext(ctx, memberID)
	if err != nil {
		fmt.Println("[InvitationRepo GetInvitationsByMemberID] error: ", err)
		return nil, err
	}

	for rows.Next() {
		var invitation model.Invitation
		if err := rows.StructScan(&invitation); err != nil {
			fmt.Println("[InvitationRepo GetInvitationsByMemberID] error: ", err)
			return nil, err
		}

		invitations = append(invitations, invitation)
	}

	return invitations, nil
}

func (r InvitationStoreRepo) GetInvitationsByGatheringID(ctx context.Context, gatheringID int64) ([]model.Invitation, error) {
	stmt, err := r.db.FPreparexContext(ctx, invitationGetByGatheringIDQuery)
	if err != nil {
		fmt.Println("[InvitationRepo GetInvitationsByGatheringID] error: ", err)
		return nil, err
	}

	defer stmt.Close()

	var invitations []model.Invitation
	rows, err := stmt.FQueryxContext(ctx, gatheringID)
	if err != nil {
		fmt.Println("[InvitationRepo GetInvitationsByMemberID] error: ", err)
		return nil, err
	}

	for rows.Next() {
		var invitation model.Invitation
		if err := rows.StructScan(&invitation); err != nil {
			fmt.Println("[InvitationRepo GetInvitationsByMemberID] error: ", err)
			return nil, err
		}

		invitations = append(invitations, invitation)
	}

	return invitations, nil
}

func (r InvitationStoreRepo) InsertInvitation(ctx context.Context, data model.Invitation) (*int64, error) {
	stmt, err := r.db.FPreparexContext(ctx, invitationInsertQuery)
	if err != nil {
		fmt.Println("[InvitationRepo InsertInvitation] error: ", err)
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, data.MemberID, data.GatheringID, data.Status)
	if err != nil {
		fmt.Println("[InvitationRepo InsertInvitation] error: ", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("[InvitationRepo InsertInvitation] error: ", err)
		return nil, err
	}

	return &id, nil
}

func (r InvitationStoreRepo) UpdateInvitation(ctx context.Context, data model.Invitation) (*int64, error) {
	stmt, err := r.db.FPreparexContext(ctx, invitationUpdateQuery)
	if err != nil {
		fmt.Println("[InvitationRepo UpdateInvitation] error: ", err)
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, data.MemberID, data.GatheringID, data.Status, data.ID)
	if err != nil {
		fmt.Println("[InvitationRepo UpdateInvitation] error: ", err)
		return nil, err
	}

	id, err := res.RowsAffected()
	if err != nil {
		fmt.Println("[InvitationRepo UpdateInvitation] error: ", err)
		return nil, err
	}

	return &id, nil
}

func (r InvitationStoreRepo) DeleteInvitation(ctx context.Context, data model.Invitation) (*int64, error) {
	stmt, err := r.db.FPreparexContext(ctx, invitationDeleteQuery)
	if err != nil {
		fmt.Println("[InvitationRepo DeleteInvitation] error: ", err)
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, data.ID)
	if err != nil {
		fmt.Println("[InvitationRepo DeleteInvitation] error: ", err)
		return nil, err
	}

	id, err := res.RowsAffected()
	if err != nil {
		fmt.Println("[InvitationRepo DeleteInvitation] error: ", err)
		return nil, err
	}

	return &id, nil
}

func (r InvitationStoreRepo) GetInvitationsByPair(ctx context.Context, gatheringID, memberID int64) (*model.Invitation, error) {
	stmt, err := r.db.FPreparexContext(ctx, invitationGetByPair)
	if err != nil {
		fmt.Println("[InvitationRepo GetInvitationsByPair] error: ", err)
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmt.Close()

	var invitation model.Invitation
	if err := stmt.FQueryRowxContext(ctx, gatheringID, memberID).StructScan(&invitation); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		fmt.Println("[InvitationRepo GetInvitationsByPair] error: ", err)
		return nil, err
	}

	return &invitation, nil
}

func NewInvitationStoreRepo(db *db.DB) repo.InvitationRepo {
	return InvitationStoreRepo{
		db: db,
	}
}
