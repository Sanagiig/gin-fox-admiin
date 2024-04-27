package system

import (
	"gin-one/global"
	"gin-one/message"
	comReq "gin-one/model/common/request"
	"gin-one/model/system"
	"gin-one/service/common"
)

type RoleService struct{}

func (service *RoleService) CreateRole(authority *system.SysRole) error {
	return global.DB.Create(authority).Error
}

func (service *RoleService) UpdateRole(authority *system.SysRole) error {
	return global.DB.Updates(authority).Error
}

func (service *RoleService) DeleteRoleByID(id string) (msgCode int, err error) {
	return common.DeleteModelByID(&system.SysRole{}, id)
}

func (service *RoleService) DeleteRoleList(ids []string) (msgCode int, err error) {
	return common.DeleteModelList(&system.SysRole{}, ids)
}

func (service *RoleService) GetRoleByID(id string) (codeMsg int, role system.SysRole, err error) {
	codeMsg, err = common.GetModelByID(&role, id)
	return
}

func (service *RoleService) GetRoleByIds(ids []string) (codeMsg int, roles []system.SysRole, err error) {
	codeMsg, err = common.GetModelByIds(&system.SysRole{}, ids, &roles)
	return
}

func (service *RoleService) GetRole(authority *system.SysRole) (int, error) {
	return common.GetModel(authority)
}

func (service *RoleService) GetRolePagination(pageInfo comReq.PageInfo, role *system.SysRole, authorityIds []string) (_ int, roles []system.SysRole, count int64, err error) {
	columns := []string{"roleName", "description"}
	vals := []string{role.RoleName, role.Description}

	if authorityIds != nil && len(authorityIds) > 0 {
		subQuery := global.DB.Model(&system.SysRoleAuthority{}).
			Select("role_id").
			Where("authority_id in ?", authorityIds)
		err = global.DB.Model(&system.SysRole{}).Where("id in (?)", subQuery).
			Scopes(
				common.AndLike(columns, vals),
			).Count(&count).Scopes(
			common.Paginate(pageInfo.Page, pageInfo.PageSize),
		).Find(&roles).Error

	} else {
		err = global.DB.Model(&system.SysRole{}).Scopes(
			common.AndLike(columns, vals),
		).Count(&count).Scopes(
			common.Paginate(pageInfo.Page, pageInfo.PageSize),
		).Find(&roles).Error
	}

	if err != nil {
		return message.QUERY_ERR, roles, count, err
	}

	return message.QUERY_OK, roles, count, err
}

var RoleServiceApp = new(RoleService)
