package system

import (
	"gin-one/global"
	"gin-one/message"
	comReq "gin-one/model/common/request"
	"gin-one/model/system"
	"gin-one/model/system/request"
	"gin-one/service/common"
)

type RoleService struct{}

func (service *RoleService) CreateRole(role *system.SysRole) error {
	return global.DB.Create(role).Error
}

func (service *RoleService) UpdateRole(role *system.SysRole) error {
	return global.DB.Updates(role).Error
}

func (service *RoleService) UpdateRoleAuthorities(req request.UpdateRoleAuthoritiesReq) (msgCode int, err error) {
	tx := global.DB.Begin()

	err = tx.Model(&system.SysRole{}).Where("id = ?", req.ID).
		Association("Authorities").Replace(req.Authorities)
	if err != nil {
		tx.Rollback()
		return message.QUERY_ERR, err
	}

	msgCode, err = AuthorizeServiceApp.UpdateRoleAuthoritiesByIds(req)
	if err != nil {
		tx.Rollback()
		return msgCode, err
	}

	tx.Commit()
	return
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

func (service *RoleService) GetCodeByIds(ids []string, res any) (codeMsg int, err error) {
	err = global.DB.Model(&system.SysRole{}).Select("id", "code").
		Where("id in (?)", ids).Find(res).Error
	if err != nil {
		return message.QUERY_ERR, err
	}

	return message.QUERY_OK, nil
}

func (service *RoleService) GetRolesByUserID(id string, res any) (codeMsg int, err error) {
	user := system.SysUser{}
	user.ID = id
	err = global.DB.Model(&user).Association("Roles").Find(res)
	if err != nil {
		return message.QUERY_ERR, err
	}
	return message.QUERY_OK, nil
}

func (service *RoleService) GetRole(role *system.SysRole) (int, error) {
	return common.GetModel(role)
}

func (service *RoleService) GetRolePagination(pageInfo comReq.PageInfo, role *system.SysRole, authorityIds []string) (_ int, roles []system.SysRole, count int64, err error) {
	columns := []string{"Name", "description"}
	vals := []string{role.Name, role.Description}

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
