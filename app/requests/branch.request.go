package requests

type CreateBranchRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=35"`
	Code        string `json:"code" validate:"required,min=3,max=15"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type UpdateBranchRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=35"`
	Code        string `json:"code" validate:"required,min=3,max=15"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
