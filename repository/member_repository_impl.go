package repository

import (
	"gorm.io/gorm"
	"library.com/data/request"
	"library.com/data/response"
)

type MemberRepositoryImpl struct {
	DB *gorm.DB
}

func NewMemberRepositoryImpl(db *gorm.DB) *MemberRepositoryImpl {
	return &MemberRepositoryImpl{
		DB: db,
	}
}

func (repo *MemberRepositoryImpl) GetAllMembers() []response.Member {
	var members []response.Member
	repo.DB.Find(&members)
	return members
}

func (repo *MemberRepositoryImpl) CreateMember(member *request.Member) error {
	db := repo.DB.Table("members").Create(&member)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (repo *MemberRepositoryImpl) ViewMemberDetail(id *int) (response.Member, error) {
	var member response.Member
	db := repo.DB.Table("members").Where("id = ?", &id).First(&member)
	if db.Error != nil {
		return response.Member{}, db.Error
	}
	return member, nil
}

func (repo *MemberRepositoryImpl) UpdateMember(member *request.Member) error {
	db := repo.DB.Table("members").Where("id = ?", &member.Id).Updates(&member)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (repo *MemberRepositoryImpl) DeleteMember(id *int) error {
	var member response.Member
	db := repo.DB.Table("members").Where("id = ?", &id).Delete(&member)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
