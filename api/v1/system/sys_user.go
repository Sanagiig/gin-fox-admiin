package system

import (
	"gin-one/message"
	comReq "gin-one/model/common/request"
	"gin-one/model/common/response"
	"gin-one/model/system"
	"gin-one/model/system/request"
	userRes "gin-one/model/system/response"
	"gin-one/utils/ctx"
	"gin-one/utils/jwt"
	"gin-one/utils/validator"
	"github.com/gin-gonic/gin"
	"time"
)

type UserApi struct{}

func (u *UserApi) Register(c *gin.Context) {
	var data request.Register
	var userModel system.SysUser
	var err error
	ip := c.ClientIP()

	if !ctx.MustBindWithCopy(c, &data, &userModel) {
		return
	}

	msgCode, ok, err := baseService.CheckCaptchaCode(ip, data.CaptchaID, data.CaptchaCode)
	if err != nil {
		response.FailWithMessage(message.CAPTCHA_ERR, err.Error(), c)
		return
	} else if !ok {
		response.FailWithMessage(msgCode, "", c)
		return
	}

	msgCode, err = userService.CreateUser(&userModel)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	} else if msgCode != message.OPER_OK {
		response.FailWithMessage(msgCode, "", c)
		return
	}

	response.OkWithDetailed(userModel, msgCode, c)
}

func (u *UserApi) Login(c *gin.Context) {
	var loginData request.Login
	err := c.ShouldBind(&loginData)
	if err != nil {
		response.FailWithMessage(message.LOGIN_ERR, err.Error(), c)
		return
	}

	err, innerErr := validator.Verify(loginData, validator.LoginVerify)
	if err != nil {
		response.FailWithMessage(message.REQ_DATA_ERR, err.Error(), c)
		return
	} else if innerErr != nil {
		response.FailWithMessage(message.LOGIN_ERR, innerErr.Error(), c)
		return
	}

	code, user, err := userService.Login(loginData)
	if err != nil {
		response.FailWithMessage(code, err.Error(), c)
		return
	}

	claims := jwt.DefaultJwtUtils.CreateClaims(request.BaseClaims{
		ID:       user.ID,
		Username: user.Username,
		NickName: user.NickName,
	})
	token, err := jwt.DefaultJwtUtils.CreateToken(claims)
	if err != nil {
		response.FailWithMessage(message.LOGIN_ERR, err.Error(), c)
		return
	}

	ctx.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Sub(time.Now()).Seconds()))
	response.OkWithDetailed(
		userRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		},
		message.LOGIN_SUCCESS,
		c,
	)
}

func (u *UserApi) CreateUser(c *gin.Context) {
	var data request.CreateUser
	var userModel system.SysUser

	if !ctx.MustBindWithCopy(c, &data, &userModel) {
		return
	}

	msgCode, err := userService.CreateUser(&userModel)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithDetailed(userModel, msgCode, c)
}

func (u *UserApi) UpdateUser(c *gin.Context) {
	var data request.UpdateUserReq
	var userModel system.SysUser

	if !ctx.MustBindWithCopy(c, &data, &userModel) {
		return
	}

	err := userService.UpdateUser(&userModel)
	if err != nil {
		response.FailWithMessage(message.OPER_DB_ERR, err.Error(), c)
		return
	}

	response.Ok(c)
}

func (u *UserApi) DeleteUserById(c *gin.Context) {
	var data comReq.GetById
	err := c.ShouldBind(&data)
	if err != nil {
		response.FailWithMessage(message.REQ_DATA_ERR, err.Error(), c)
		return
	}

	code, err := userService.DeleteUserById(data.ID)
	if err != nil {
		response.FailWithMessage(message.OPER_ERR, err.Error(), c)
		return
	}
	response.OkWithMessage(code, c)
}

func (u *UserApi) DeleteUserList(c *gin.Context) {
	var data comReq.GetByIds
	err := c.ShouldBind(&data)
	if err != nil {
		response.FailWithMessage(message.REQ_DATA_ERR, err.Error(), c)
		return
	}

	code, err := userService.DeleteUserList(data.Ids)
	if err != nil {
		response.FailWithMessage(message.OPER_ERR, err.Error(), c)
		return
	}
	response.OkWithMessage(code, c)
}

func (u *UserApi) GetUserByID(c *gin.Context) {
	var data comReq.GetById

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, role, err := userService.GetUserByID(data.ID)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithDetailed(role, msgCode, c)
}

func (u *UserApi) GetUser(c *gin.Context) {
	var data request.GetUserReq
	var userModel system.SysUser

	if !ctx.MustBindWithCopy(c, &data, &userModel) {
		return
	}

	msgCode, err := userService.GetUser(&userModel)

	if err != nil {
		response.FailWithMessage(message.OPER_ERR, err.Error(), c)
		return
	}
	response.OkWithDetailed(userModel, msgCode, c)
}

func (u *UserApi) GetUserPagination(c *gin.Context) {
	var data request.GetUserPaginationReq
	var userModel system.SysUser

	if !ctx.MustBindWithCopy(c, &data, &userModel) {
		return
	}

	msgCode, users, count, err := userService.GetUserPagination(data.PageInfo, &userModel, data.RoleIds)
	if err != nil {
		response.FailWithPage(response.WrapPageData(data.PageInfo, count, users), msgCode, err.Error(), c)
		return
	}
	response.OkWithDetailed(response.WrapPageData(data.PageInfo, count, users), msgCode, c)
}
