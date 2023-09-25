package requests

type CreateMemberOccupationRequest struct {
	Company    string `json:"company" validate:"required,max=50"`
	Department string `json:"department" validate:"required,max=50"`
	Address    string `json:"address" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required,max=6"`
	Phone      string `json:"phone" validate:"phoneNumberOrEmpty"`
	Fax        string `json:"fax" validate:"phoneNumberOrEmpty"`
	Email      string `json:"email" validate:"emailOrEmpty"`
}

type UpdateMemberOccupationRequest struct {
	Company    string `json:"company" validate:"required,max=50"`
	Department string `json:"department" validate:"required,max=50"`
	Address    string `json:"address" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required,max=6"`
	Phone      string `json:"phone" validate:"phoneNumberOrEmpty"`
	Fax        string `json:"fax" validate:"phoneNumberOrEmpty"`
	Email      string `json:"email" validate:"emailOrEmpty"`
}
