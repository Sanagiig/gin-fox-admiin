package system

import (
	"errors"
	"gin-one/global"
	"gin-one/message"
	"gin-one/model/common"
	comReq "gin-one/model/common/request"
	"gin-one/model/system"
	"gin-one/model/system/request"
	comService "gin-one/service/common"
	"gin-one/utils"
	"gorm.io/gorm"
)

type UserService struct{}

func (service *UserService) Login(login request.Login) (code int, user system.SysUser, err error) {
	user.Username = login.Username
	err = global.DB.Model(&user).Preload("Roles").First(&user).Error
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

	isExist, err = comService.IsModelDataExist(system.SysUser{Username: u.Username})
	if err != nil {
		return message.OPER_DB_ERR, err
	} else if isExist {
		return message.OPER_FAILED, errors.New(global.Msg.Msg(message.USER_IS_EXIST))
	}

	if u.Roles == nil {
		u.Roles = []system.SysRole{
			system.SysRole{
				UUIDModel: common.UUIDModel{
					ID: "1",
				},
			},
		}
	}

	u.Password = utils.BcryptHash(u.Password)

	err = global.DB.Create(u).Error
	if err != nil {
		return message.OPER_DB_ERR, err
	}

	roleIds := make([]string, len(u.Roles))
	for i, _ := range roleIds {
		roleIds[i] = u.Roles[i].ID
	}
	msgCode, u.Roles, err = RoleServiceApp.GetRoleByIds(roleIds)
	if err != nil {
		return
	}
	return message.OPER_OK, nil
}

func (service *UserService) UpdateUser(u *system.SysUser) (err error) {
	return global.DB.Model(u).Updates(u).Error
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
