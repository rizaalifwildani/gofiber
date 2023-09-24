package responses

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MemberResponse struct {
	ID             uuid.UUID        `json:"id"`
	Phone          string           `json:"phone"`
	Email          string           `json:"email"`
	FirstName      string           `json:"firstName"`
	LastName       string           `json:"lastName"`
	IdentityNumber string           `json:"identityNumber"`
	PlaceOfBirth   string           `json:"placeOfBirth"`
	Birthdate      string           `json:"birthdate"`
	Gender         string           `json:"gender"`
	Nationality    string           `json:"nationality"`
	Address        string           `json:"address"`
	PostalCode     string           `json:"postalCode"`
	HomePhone      string           `json:"homePhone"`
	OfficePhone    string           `json:"officePhone"`
	Education      string           `json:"education"`
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

func NewMemberCollections(ctx *fiber.Ctx, m []models.Member) error {
	data := []MemberResponse{}

	for _, v := range m {
		data = append(data, MemberResponse{
			ID:             v.ID,
			Phone:          v.Phone,
			Email:          v.Email,
			FirstName:      v.FirstName,
			LastName:       v.LastName,
			IdentityNumber: v.IdentityNumber,
		})
	}

	return SuccessResponse(ctx, data)
}
