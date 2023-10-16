package service

import (
	"context"
	"fmt"

	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/pkg/dto"
	errors "github.com/uchupx/geek-garden-test/pkg/error"
	"github.com/uchupx/geek-garden-test/pkg/helper"
)

type GatheringService struct {
	GatheringRepo repo.GatheringRepo
}

func (s GatheringService) GetGatherings(ctx context.Context) ([]dto.Gathering, error) {
	var dtos []dto.Gathering

	items, err := s.GatheringRepo.GetGatherings(ctx)
	if err != nil {
		fmt.Println("[GatheringService GetGatherings] error: ", err)
		return nil, err
	}

	for _, item := range items {
		var gatheringDTO dto.Gathering

		gatheringDTO.FromModel(item)
		dtos = append(dtos, gatheringDTO)
	}

	return dtos, nil
}

func (s GatheringService) InsertGathering(ctx context.Context, data dto.Gathering) (*int64, error) {
	return s.GatheringRepo.InsertGathering(ctx, data.ToModel())
}

func (s GatheringService) GetGatheringByID(ctx context.Context, id int64) (*dto.Gathering, error) {
	m, err := s.GatheringRepo.GetGatheringByID(ctx, id)
	if err != nil {
		return nil, err
	}

	gathering := dto.Gathering{}
	gathering.FromModel(*m)

	return &gathering, nil
}

func (s GatheringService) UpdateGathering(ctx context.Context, id int64, data dto.GatheringPutV1) (*int64, error) {
	gathering, err := s.GatheringRepo.GetGatheringByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if gathering == nil {
		return nil, errors.ErrDataNotFound
	}

	gatheringDto := dto.Gathering{}
	gatheringDto.FromModel(*gathering)

	if data.Creator != nil {
		gatheringDto.Creator = *data.Creator
	}

	if data.Location != nil {
		gatheringDto.Location = *data.Location
	}

	if data.ScheduledAt != nil {
		parseTime, err := helper.ParseTime(*data.ScheduledAt)
		if err != nil {
			return nil, err
		}
		gatheringDto.ScheduledAt = *parseTime
	}

	if data.Name != nil {
		gatheringDto.Name = *data.Name
	}

	if data.Type != nil {
		gatheringDto.Type = *data.Type
	}

	return s.GatheringRepo.UpdateGathering(ctx, gatheringDto.ToModel())
}

func (s GatheringService) DeleteGathering(ctx context.Context, id int64) error {
	data, err := s.GatheringRepo.GetGatheringByID(ctx, id)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.ErrDataNotFound
	}

	_, err = s.GatheringRepo.DeleteGathering(ctx, *data)
	if err != nil {
		return err
	}

	return nil
}
