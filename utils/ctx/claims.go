package ctx

import (
	"gin-one/global"
	"gin-one/model/system/request"
	"gin-one/utils/jwt"
	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	token := GetToken(c)
	j := jwt.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.Log.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.Username
	}
}
