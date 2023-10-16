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
	gatheringGetByIDQuery = `SELECT * FROM gatherings WHERE id = ?`
	gatheringGetAllQuery  = `SELECT * FROM gatherings`
	gatheringInsertQuery  = `INSERT INTO gatherings (creator, location, scheduled_at, name, type) VALUES (?, ?, ?, ?, ?)`
	gatheringUpdateQuery  = `UPDATE gatherings SET creator = ?, location = ?, scheduled_at = ?, name = ?, type = ? WHERE id = ?`
	gatheringDeleteQuery  = `DELETE FROM gatherings WHERE id = ?`
)

type GatheringStoreRepo struct {
	db *db.DB
}

func NewGatheringStore(db *db.DB) repo.GatheringRepo {
	return &GatheringStoreRepo{
		db: db,
	}
}

func (r GatheringStoreRepo) GetGatheringByID(ctx context.Context, id int64) (*model.Gathering, error) {
	stmt, err := r.db.FPreparexContext(ctx, gatheringGetByIDQuery)
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo GetGatheringByID] error preparing query: %v", err)
	}

	defer stmt.Close()

	var gathering model.Gathering
	if err := stmt.FQueryRowxContext(ctx, id).StructScan(&gathering); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo GetGatheringByID] error executing query: %v", err)
	}

	return &gathering, nil
}
func (r GatheringStoreRepo) GetGatherings(ctx context.Context) ([]model.Gathering, error) {
	stmt, err := r.db.FPreparexContext(ctx, gatheringGetAllQuery)
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo GetGatherings] error preparing query: %v", err)
	}

	defer stmt.Close()

	var gatherings []model.Gathering
	rows, err := stmt.FQueryxContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo GetGatherings] error executing query: %v", err)
	}

	for rows.Next() {
		var gathering model.Gathering
		if err := rows.StructScan(&gathering); err != nil {
			return nil, fmt.Errorf("[GatheringStoreRepo GetGatherings] error scanning row: %v", err)
		}

		gatherings = append(gatherings, gathering)
	}

	return gatherings, nil
}
func (r GatheringStoreRepo) InsertGathering(ctx context.Context, data model.Gathering) (*int64, error) {
	stmt, err := r.db.FPreparexContext(ctx, gatheringInsertQuery)
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo InsertGathering] error preparing query: %v", err)
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, data.Creator, data.Location, data.ScheduledAt, data.Name, data.Type)
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo InsertGathering] error executing query: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo InsertGathering] error getting last insert id: %v", err)
	}

	return &id, nil
}

func (r GatheringStoreRepo) UpdateGathering(ctx context.Context, data model.Gathering) (*int64, error) {
	stmt, err := r.db.FPreparexContext(ctx, gatheringUpdateQuery)
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo UpdateGathering] error preparing query: %v", err)
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, data.Creator, data.Location, data.ScheduledAt, data.Name, data.Type, data.ID)
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo UpdateGathering] error executing query: %v", err)
	}

	id, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo UpdateGathering] error getting rows affected: %v", err)
	}

	return &id, nil
}

func (r GatheringStoreRepo) DeleteGathering(ctx context.Context, data model.Gathering) (*int64, error) {
	stmt, err := r.db.FPreparexContext(ctx, gatheringDeleteQuery)
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo DeleteGathering] error preparing query: %v", err)
	}

	defer stmt.Close()

	res, err := stmt.FExecContext(ctx, data.ID)
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo DeleteGathering] error executing query: %v", err)
	}

	id, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("[GatheringStoreRepo DeleteGathering] error getting rows affected: %v", err)
	}

	return &id, nil
}
