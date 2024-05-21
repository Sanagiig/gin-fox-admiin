package system

import (
	"gin-one/global"
	"gin-one/message"
	comReq "gin-one/model/common/request"
	"gin-one/model/system"
	"gin-one/service/common"
	"gorm.io/gorm"
)

type AuthorityService struct{}

func (service *AuthorityService) CreateAuthority(authority *system.SysAuthority) (int, error) {
	ok, err := common.IsModelDataExist(system.SysAuthority{
		AuthorityVal: authority.AuthorityVal,
	})

	if err != nil {
		return message.OPER_ERR, err
	} else if ok {
		return message.DATA_EXIST, nil
	}

	err = global.DB.Create(authority).Error
	if err != nil {
		return message.OPER_ERR, err
	}

	return message.OPER_OK, err
}

func (service *AuthorityService) UpdateAuthority(authority *system.SysAuthority) error {
	return global.DB.Updates(authority).Error
}

func (service *AuthorityService) DeleteAuthorityByID(id string) (int, error) {
	return service.DeleteAuthorityList([]string{id})
}

func (service *AuthorityService) DeleteAuthorityList(ids []string) (msgCode int, err error) {
	childrenIds := make([]string, 0, len(ids))
	msgCode, err = service.GetChildrenIdsByParentIds(ids, &childrenIds)
	if err != nil {
		return
	}

	childrenIds = append(childrenIds, ids...)
	return common.DeleteModelList(&system.SysAuthority{}, childrenIds)
}

func (service *AuthorityService) GetChildrenIdsByParentIds(ids []string, res *[]string) (msgCode int, err error) {
	children := make([]map[string]any, 0)
	msgCode, err = common.GetModelChildrenByParentIds(&system.SysAuthority{}, ids, &children, "id")
	if err != nil {
		return msgCode, err
	}

	for _, child := range children {
		*res = append(*res, child["id"].(string))
	}
	return
}

func (service *AuthorityService) GetAuthorityByID(id string) (msgCode int, auth system.SysAuthority, err error) {
	msgCode, err = common.GetModelByID(&auth, id)
	return
}

func (service *AuthorityService) GetAuthorityCodeByIds(ids []string, res any) (msgCode int, err error) {
	err = global.DB.Model(&system.SysAuthority{}).Select("id", "code").Where("id in ?", ids).Find(&res).Error
	if err != nil {
		return message.QUERY_ERR, err
	}
	return message.QUERY_OK, nil
}

func (service *AuthorityService) GetAuthority(authority *system.SysAuthority) (int, error) {
	return common.GetModel(authority)
}

func (service *AuthorityService) GetServiceDb(auth *system.SysAuthority, roleIds []string) *gorm.DB {
	columns := []string{"authority_name", "authority_val", "description"}
	vals := []string{auth.AuthorityName, auth.AuthorityVal, auth.Description}
	db := global.DB.Model(&system.SysAuthority{})

	if auth.ParentID != "" {
		db.Where("parent_id = ? ", auth.ParentID)
	}

	if auth.Status != "" {
		db = db.Where("status = ?", auth.Status)
	}

	if auth.AuthorityType != "" {
		db = db.Where("authority_type = ?", auth.AuthorityType)
	}

	if roleIds != nil && len(roleIds) > 0 {
		subQuery := global.DB.Model(&system.SysRoleAuthority{}).
			Select("authority_id").
			Where("role_id in ?", roleIds)
		db.Where("id in (?)", subQuery)
	}

	db.Scopes(common.AndLike(columns, vals))

	return db
}

func (service *AuthorityService) GetAuthorityPagination(pageInfo comReq.PageInfo, auth *system.SysAuthority, roleIds []string) (msgCode int, auths []*system.SysAuthority, count int64, err error) {
	db := service.GetServiceDb(auth, roleIds)

	if auth.ParentID == "" {
		db.Where("parent_id IS NULL OR parent_id = ''")
	}

	err = db.Count(&count).Scopes(common.Paginate(pageInfo.Page, pageInfo.PageSize)).Find(&auths).Error
	if err != nil {
		return message.QUERY_ERR, auths, count, err
	}

	parentIds := make([]string, 0, len(auths))
	for i, _ := range auths {
		parentIds = append(parentIds, auths[i].ID)
	}

	msgCode, auths, err = service.GetAuthorityTree(&system.SysAuthority{}, parentIds)
	return
}

func (service *AuthorityService) GetAuthorityChildrenByParentIds(parentIds []string, auth *system.SysAuthority) (msgCode int, auths []system.SysAuthority, err error) {
	db := service.GetServiceDb(auth, nil)

	msgCode, err = common.GetChildrenByParents(db, parentIds, &auths)
	return msgCode, auths, err
}

func (service *AuthorityService) GetAuthorityTree(auth *system.SysAuthority, parentIds []string) (msgCode int, auths []*system.SysAuthority, err error) {
	var res []system.SysAuthority
	var totalRes []*system.SysAuthority
	var tmpParentIds []string

	if parentIds == nil || len(parentIds) == 0 {
		parentIds = append(parentIds, auth.ParentID)
	}

	err = global.DB.Model(&system.SysAuthority{}).Where("id in ?", parentIds).Find(&totalRes).Error
	if err != nil {
		return message.QUERY_ERR, nil, err
	}

	tmpParentIds = append(tmpParentIds, parentIds...)
	for {
		msgCode, res, err = service.GetAuthorityChildrenByParentIds(tmpParentIds, auth)
		if err != nil {
			return msgCode, auths, err
		}
		if len(res) == 0 {
			break
		}

		tmpParentIds = make([]string, 0, len(res))
		for i, _ := range res {
			tmpParentIds = append(tmpParentIds, res[i].ID)
			totalRes = append(totalRes, &res[i])
		}
	}

	auths, err = common.AssembleTree(totalRes, parentIds)
	return message.QUERY_OK, auths, nil
}

var AuthorityServiceApp = new(AuthorityService)
