package repository

import (
	"library.com/data/request"
	"library.com/data/response"
)

type MemberRepository interface {
	GetAllMembers() []response.Member
	CreateMember(member *request.Member) error
	ViewMemberDetail(id *int) (response.Member, error)
	UpdateMember(member *request.Member) error
	DeleteMember(id *int) error
}
