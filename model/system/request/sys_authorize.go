package request

type AuthorizeData struct {
	ID   string `json:"id" binding:"required" example:"id"`
	Code string `json:"code" binding:"required" example:"code"`
}

type UpdateUserRolesData struct {
	ID    string              `json:"id" binding:"required" example:"id"`
	Roles []map[string]string `json:"roles" binding:"required" example:"role"`
}

type UpdateRoleAuthoritiesReq struct {
	ID          string   `json:"id" binding:"required" example:"id"`
	Authorities []string `json:"authorities" binding:"required" example:"roles"`
}

type UpdateRoleAuthoritiesData struct {
	ID          string              `json:"id" binding:"required" example:"id"`
	Authorities []map[string]string `json:"roles" binding:"required" example:"roles"`
}
