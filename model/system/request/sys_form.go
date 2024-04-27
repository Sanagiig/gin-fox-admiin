package request

type CreateForm struct {
	User     string `json:"user" form:"user" example:"用户名"`
	Password string `json:"password" form:"password" example:"密码"`
	U1       string `uri:"u1"`
}

type GetForm struct {
	ID string `form:"id" json:"id" binding:"required"`
}
