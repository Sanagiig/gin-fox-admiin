package system

import (
	"gin-one/global"
	"gin-one/message"
	comReq "gin-one/model/common/request"
	"gin-one/model/system"
	"gin-one/service/common"
)

type AuthorityService struct{}

func (service *AuthorityService) CreateAuthority(authority *system.SysAuthority) error {
	return global.DB.Create(authority).Error
}

func (service *AuthorityService) UpdateAuthority(authority *system.SysAuthority) error {
	return global.DB.Updates(authority).Error
}

func (service *AuthorityService) DeleteAuthorityByID(id string) (int, error) {
	return common.DeleteModelByID(&system.SysAuthority{}, id)
}

func (service *AuthorityService) DeleteAuthorityList(ids []string) (int, error) {
	return common.DeleteModelList(&system.SysAuthority{}, ids)
}

func (service *AuthorityService) GetAuthorityByID(id string) (codeMsg int, auth system.SysAuthority, err error) {
	codeMsg, err = common.GetModelByID(&auth, id)
	return
}

func (service *AuthorityService) GetAuthority(authority *system.SysAuthority) (int, error) {
	return common.GetModel(authority)
}

func (service *AuthorityService) GetAuthorityPagination(pageInfo comReq.PageInfo, auth *system.SysAuthority, roleIds []string) (msgCode int, auths []system.SysAuthority, count int64, err error) {
	columns := []string{"authorityName", "authorityVal", "description"}
	vals := []string{auth.AuthorityName, auth.AuthorityVal, auth.Description}

	if roleIds != nil && len(roleIds) > 0 {
		subQuery := global.DB.Model(&system.SysRoleAuthority{}).
			Select("authority_id").
			Where("role_id in ?", roleIds)

		err = global.DB.Model(&system.SysAuthority{}).Where("id in (?)", subQuery).
			Scopes(
				common.AndLike(columns, vals),
			).Count(&count).Scopes(
			common.Paginate(pageInfo.Page, pageInfo.PageSize),
		).Find(&auths).Error

	} else {
		err = global.DB.Model(&system.SysAuthority{}).Model(&system.SysAuthority{}).Scopes(
			common.AndLike(columns, vals),
		).Count(&count).Scopes(
			common.Paginate(pageInfo.Page, pageInfo.PageSize),
		).Find(&auths).Error
	}

	if err != nil {
		return message.QUERY_ERR, auths, count, err
	}

	return message.QUERY_OK, auths, count, err
}

var AuthorityServiceApp = new(AuthorityService)
