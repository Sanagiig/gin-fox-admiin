package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// GetById Find by id structure
type GetById struct {
	ID string `json:"id" form:"id"` // 主键ID
}

type GetByIds struct {
	Ids []string `json:"ids" form:"ids"`
}
