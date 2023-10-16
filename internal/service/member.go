package service

import (
	"context"
	"fmt"

	"github.com/uchupx/geek-garden-test/internal/repo"
	"github.com/uchupx/geek-garden-test/pkg/dto"
)

type MemberService struct {
	MemberRepo repo.MemberRepo
}

func (s MemberService) GetMembers(ctx context.Context) ([]dto.Member, error) {
	var dtos []dto.Member

	items, err := s.MemberRepo.GetMembers(ctx)
	if err != nil {
		fmt.Println("[MemberService GetMembers] error: ", err)
		return nil, err
	}

	for _, item := range items {
		var memberDTO dto.Member

		memberDTO.FromModel(item)
		dtos = append(dtos, memberDTO)
	}

	return dtos, nil
}

func (s MemberService) InsertMember(ctx context.Context, data dto.Member) (*int64, error) {
	return s.MemberRepo.InsertMember(ctx, data.ToModel())
}

func (s MemberService) GetMemberByID(ctx context.Context, id int64) (*dto.Member, error) {
	m, err := s.MemberRepo.GetMemberByID(ctx, id)
	if err != nil {
		return nil, err
	}

	member := dto.Member{}
	member.FromModel(*m)

	return &member, nil
}

func (s MemberService) UpdateMember(ctx context.Context, id int64, data dto.MemberPutV1) (*int64, error) {
	member, err := s.MemberRepo.GetMemberByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if member == nil {
		return nil, fmt.Errorf("data not found")
	}

	memberDto := dto.Member{}
	memberDto.FromModel(*member)

	if data.FirstName != nil {
		memberDto.FirstName = *data.FirstName
	}

	if data.LastName != nil {
		memberDto.LastName = data.LastName
	}

	if data.Email != nil {
		memberDto.Email = *data.Email
	}

	return s.MemberRepo.UpdateMember(ctx, memberDto.ToModel())
}

func (s MemberService) DeleteMember(ctx context.Context, id int64) error {
	data, err := s.MemberRepo.GetMemberByID(ctx, id)
	if err != nil {
		return err
	}

	if data == nil {
		return fmt.Errorf("data not found")
	}

	_, err = s.MemberRepo.DeleteMember(ctx, *data)
	if err != nil {
		return err
	}

	return nil
}
