package service

import (
	"library.com/data/request"
	"library.com/data/response"
)

type MemberService interface {
	FetchAllMembers() []response.Member
	AddMember(member *request.Member) error
	GetMemberDetailById(id *int) (response.Member, error)
	UpdateMember(member *request.Member) error
	DeleteMember(id *int) error
}
