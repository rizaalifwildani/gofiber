package requests

type CreatePermissionRequest struct {
	DisplayName string `json:"displayName" validate:"required,min=3,max=35"`
}

type UpdatePermissionRequest struct {
	DisplayName string `json:"displayName" validate:"required,min=3,max=35"`
}
