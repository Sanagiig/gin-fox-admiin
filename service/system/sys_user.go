package system

import (
	"errors"
	"gin-one/global"
	"gin-one/message"
	comReq "gin-one/model/common/request"
	"gin-one/model/system"
	"gin-one/model/system/request"
	comService "gin-one/service/common"
	"gin-one/utils"
	"gin-one/utils/process_line"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserService struct{}

func (service *UserService) Login(login request.Login) (code int, user system.SysUser, err error) {
	user.Username = login.Username

	err = global.DB.Preload("Roles").Where("username = ?", login.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return message.USERNAME_OR_PASS_FAILED, user, err
		}
		return message.LOGIN_ERR, user, err
	}

	if !utils.BcryptCheck(login.Password, user.Password) {
		return message.USERNAME_OR_PASS_FAILED, user, err
	}

	return message.LOGIN_SUCCESS, user, err
}

func (service *UserService) CreateUser(u *system.SysUser) (msgCode int, err error) {
	var isExist bool
	var userRolesData request.UpdateUserRolesData
	isExist, err = comService.IsModelDataExist(system.SysUser{Username: u.Username})
	if err != nil {
		return message.OPER_DB_ERR, err
	} else if isExist {
		return message.USER_IS_EXIST, err
	}

	if u.Roles == nil {
		var r system.SysRole
		r.ID = "1"
		u.Roles = []system.SysRole{r}
	}

	u.Password = utils.BcryptHash(u.Password)

	tx := global.DB.Begin()
	pl := process_line.New()

	pl.Then(func() (int, error) {
		return 0, tx.Create(&u).Error
	}).Then(func() (int, error) {
		roleIds := make([]string, len(u.Roles))
		for i, _ := range roleIds {
			roleIds[i] = u.Roles[i].ID
		}

		msgCode, u.Roles, err = RoleServiceApp.GetRoleByIds(roleIds)
		return msgCode, err
	}).Then(func() (int, error) {
		err = copier.Copy(&userRolesData, &u.Roles)
		return 0, err
	}).Then(func() (int, error) {
		updateUserRolesReq := request.UpdateUserRolesReq{
			RoleIds: make([]string, 0, len(u.Roles)),
		}

		for _, role := range u.Roles {
			updateUserRolesReq.RoleIds = append(updateUserRolesReq.RoleIds, role.ID)
		}

		userRolesData.ID = u.ID
		return service.UpdateUserRoles(updateUserRolesReq)
	}).Catch(func(e error) (isCatchContinue bool) {
		tx.Rollback()
		return false
	})

	tx.Commit()
	return message.OPER_OK, nil
}

func (service *UserService) UpdateUser(u *system.SysUser) (err error) {
	return global.DB.Model(u).Updates(u).Error
}

func (service *UserService) AddUserRoles(req request.UpdateUserRolesReq) (msgCode int, err error) {
	roleCodes := make([]map[string]string, 0, len(req.RoleIds))

	msgCode, err = RoleServiceApp.GetCodeByIds(req.RoleIds, &roleCodes)
	if err != nil {
		return message.OPER_ERR, err
	} else if len(roleCodes) != len(req.RoleIds) {
		return message.SOME_ROLES_NOT_EXIST, nil
	}

	u := system.SysUser{}
	u.ID = req.ID
	roles := make([]system.SysRole, 0, len(req.RoleIds))
	for _, roleId := range req.RoleIds {
		r := system.SysRole{}
		r.ID = roleId
		roles = append(roles, r)
	}

	tx := global.DB.Model(&u).Begin()
	err = tx.Association("Roles").Append(roles)
	if err != nil {
		return message.OPER_ERR, err
	}

	codes := make([]string, 0, len(req.RoleIds))
	for _, role := range roleCodes {
		codes = append(codes, role["code"])
	}

	msgCode, err = AuthorizeServiceApp.AddUserRoles(req.ID, codes)
	if err != nil {
		tx.Rollback()
		return message.OPER_ERR, err
	}
	tx.Commit()
	return message.OPER_OK, nil
}

func (service *UserService) UpdateUserRoles(req request.UpdateUserRolesReq) (msgCode int, err error) {
	u := system.SysUser{}
	roles := make([]system.SysRole, 0, len(req.RoleIds))

	for _, roleId := range req.RoleIds {
		r := system.SysRole{}
		r.ID = roleId
		roles = append(roles, r)
	}

	u.ID = req.ID
	tx := global.DB.Model(&u).Begin()
	err = tx.Association("Roles").Clear()
	if err != nil {
		return message.OPER_DB_ERR, err
	}

	err = tx.Association("Roles").Append(roles)
	if err != nil {
		tx.Rollback()
		return message.OPER_DB_ERR, err
	}

	msgCode, err = AuthorizeServiceApp.UpdateUserRolesByIds(req)
	if err != nil {
		tx.Rollback()
	}

	tx.Commit()
	return
}

func (service *UserService) DeleteUserById(id string) (int, error) {
	return comService.DeleteModelByID(&system.SysUser{}, id)
}

func (service *UserService) DeleteUserList(ids []string) (int, error) {
	return comService.DeleteModelList(&system.SysUser{}, ids)
}

func (service *UserService) GetUserByID(id string) (msgCode int, user system.SysUser, err error) {
	msgCode, err = comService.GetModelByID(&user, id)
	return
}

func (service *UserService) GetUser(u *system.SysUser) (msgCode int, err error) {
	return comService.GetModel(u)
}

func (service *UserService) GetUserPagination(pageInfo comReq.PageInfo, u *system.SysUser, roleIds []string) (msgCode int, users []system.SysUser, count int64, err error) {
	columns := []string{"username", "nick_name", "phone", "email"}
	vals := []string{u.Username, u.NickName, u.Phone, u.Email}

	if roleIds != nil && len(roleIds) > 0 {
		subQuery := global.DB.Model(&system.SysUserRole{}).
			Select("user_id").
			Where("role_id in ?", roleIds)

		err = global.DB.Model(&system.SysUser{}).Where("id in (?)", subQuery).
			Scopes(
				comService.AndLike(columns, vals),
			).Count(&count).Scopes(
			comService.Paginate(pageInfo.Page, pageInfo.PageSize),
		).Find(&users).Error

	} else {
		err = global.DB.Model(&system.SysUser{}).Scopes(
			comService.AndLike(columns, vals),
		).Count(&count).Scopes(
			comService.Paginate(pageInfo.Page, pageInfo.PageSize),
		).Find(&users).Error
	}

	if err != nil {
		return message.QUERY_ERR, users, count, err
	}

	return message.QUERY_OK, users, count, err
}

var UserServiceApp = new(UserService)
