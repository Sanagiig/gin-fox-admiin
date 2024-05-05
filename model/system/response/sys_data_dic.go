package response

import "gin-one/model/system"

type DataDicRes struct {
	ParentName string `json:"parentName"`
	system.SysDataDic
}
