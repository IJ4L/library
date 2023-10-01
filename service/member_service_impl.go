package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"library.com/data/request"
	"library.com/data/response"
	"library.com/repository"
	"library.com/utils"
)

type MemberServiceImpl struct {
	memberRepository repository.MemberRepository
	validate         *validator.Validate
}

func NewMemberServiceImpl(memberRepository repository.MemberRepository, validator *validator.Validate) *MemberServiceImpl {
	return &MemberServiceImpl{
		memberRepository: memberRepository,
		validate:         validator,
	}
}

func (service *MemberServiceImpl) FetchAllMembers() []response.Member {
	return service.memberRepository.GetAllMembers()
}

func (service *MemberServiceImpl) AddMember(member *request.Member) error {
	if err := utils.ValidateFunc(&member, service.validate); err != nil {
		return err
	}

	if err := service.memberRepository.CreateMember(member); err != nil {
		return err
	}

	return nil
}

func (service *MemberServiceImpl) GetMemberDetailById(id *int) (response.Member, error) {
	if *id < 0 {
		return response.Member{}, fmt.Errorf("error: ID cannot be nil")
	}

	member, err := service.memberRepository.ViewMemberDetail(id)
	if err != nil {
		return response.Member{}, err
	}

	return member, nil
}

func (service *MemberServiceImpl) UpdateMember(member *request.Member) error {
	if err := utils.ValidateFunc(&member, service.validate); err != nil {
		return err
	}

	if err := service.memberRepository.UpdateMember(member); err != nil {
		return err
	}

	return nil
}

func (service *MemberServiceImpl) DeleteMember(id *int) error {
	if *id < 0 {
		return fmt.Errorf("error: ID cannot be nil")
	}

	if err := service.memberRepository.DeleteMember(id); err != nil {
		return err
	}

	return nil
}
