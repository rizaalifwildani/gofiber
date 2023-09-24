package responses

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/morkid/paginate"
)

type MemberResponse struct {
	ID             uuid.UUID        `json:"id"`
	Phone          string           `json:"phone"`
	Email          string           `json:"email"`
	FirstName      string           `json:"firstName"`
	LastName       string           `json:"lastName"`
	IdentityNumber string           `json:"identityNumber,omitempty"`
	PlaceOfBirth   string           `json:"placeOfBirth,omitempty"`
	Birthdate      string           `json:"birthdate,omitempty"`
	Gender         string           `json:"gender,omitempty"`
	Nationality    string           `json:"nationality,omitempty"`
	Address        string           `json:"address,omitempty"`
	PostalCode     string           `json:"postalCode,omitempty"`
	HomePhone      string           `json:"homePhone,omitempty"`
	OfficePhone    string           `json:"officePhone,omitempty"`
	Education      string           `json:"education,omitempty"`
	Branches       []BranchResponse `json:"branches,omitempty"`
}

func NewMemberResponse(ctx *fiber.Ctx, m models.Member) error {
	branches := []BranchResponse{}
	for _, branch := range m.Branches {
		branches = append(branches, BranchResponse{
			ID:          branch.BranchID,
			Name:        branch.Branch.Name,
			Code:        branch.Branch.Code,
			Address:     branch.Branch.Address,
			Description: branch.Branch.Description,
			Status:      branch.Status,
		})
	}
	data := MemberResponse{
		ID:             m.ID,
		Phone:          m.Phone,
		Email:          m.Email,
		FirstName:      m.FirstName,
		LastName:       m.LastName,
		IdentityNumber: m.IdentityNumber,
		PlaceOfBirth:   m.PlaceOfBirth,
		Birthdate:      m.Birthdate.Format("2006-01-02"),
		Gender:         m.Gender,
		Nationality:    m.Nationality,
		Address:        m.Address,
		PostalCode:     m.PostalCode,
		HomePhone:      m.HomePhone,
		OfficePhone:    m.OfficePhone,
		Education:      m.Education,
		Branches:       branches,
	}
	return SuccessResponse(ctx, data)
}

func NewMemberCollections(ctx *fiber.Ctx, data paginate.Page) error {
	members := data.Items.(*[]models.Member)
	memberResponses := []MemberResponse{}

	for _, v := range *members {
		memberResponses = append(memberResponses, MemberResponse{
			ID:             v.ID,
			Phone:          v.Phone,
			Email:          v.Email,
			FirstName:      v.FirstName,
			LastName:       v.LastName,
			IdentityNumber: v.IdentityNumber,
		})
	}
	data.Items = memberResponses

	return PaginationResponse(ctx, data)
}
