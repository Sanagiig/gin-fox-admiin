package system

import (
	"gin-one/message"
	"gin-one/model/system"
	"gin-one/model/system/request"
	"gin-one/service/common"
)

type AuthorizeService struct{}

func (service *AuthorizeService) UpdateUserRolesByIds(req request.UpdateUserRolesReq) (msgCode int, err error) {
	roles := make([]map[string]string, 0, len(req.Roles))

	msgCode, err = RoleServiceApp.GetCodeByIds(req.Roles, &roles)
	if err != nil {
		return
	}

	codes := make([]string, 0, len(roles))
	for _, role := range roles {
		codes = append(codes, role["code"])
	}

	return service.UpdateUserRoles(req.ID, codes)
}

func (service *AuthorizeService) UpdateUserRoles(id string, codes []string) (msgCode int, err error) {
	ok, err := CasbinServiceApp.Casbin().AddRolesForUser(id, codes)
	if err != nil {
		return message.OPER_ERR, err
	} else if !ok {
		return message.OPER_FAILED, err
	}

	return
}

func (service *AuthorizeService) UpdateRoleAuthoritiesByIds(req request.UpdateRoleAuthoritiesReq) (msgCode int, err error) {
	auths := make([]map[string]string, 0, len(req.Authorities))

	msgCode, err = common.GetModelByIds(&system.SysAuthority{}, req.Authorities, &auths)
	if err != nil {
		return
	}

	codes := make([]string, 0, len(req.Authorities))
	for _, auth := range auths {
		codes = append(codes, auth["code"])
	}

	return service.UpdateUserRoles(req.ID, codes)
}

func (service *AuthorizeService) UpdateRoleAuthorities(id string, authsData []map[string]string) (msgCode int, err error) {
	ok, err := CasbinServiceApp.Casbin().RemoveGroupingPolicy(id)
	if err != nil {
		return message.OPER_ERR, err
	} else if !ok {
		return message.OPER_FAILED, err
	}

	rules := make([][]string, 0, len(authsData))
	for _, auth := range authsData {
		rules = append(rules, []string{
			id,
			auth["authority_val"],
			auth["authority_type"],
		})
	}
	ok, err = CasbinServiceApp.Casbin().AddPolicies(rules)
	if err != nil {
		return message.OPER_ERR, err
	} else if !ok {
		return message.OPER_FAILED, err
	}
	return
}

var AuthorizeServiceApp = new(AuthorizeService)
