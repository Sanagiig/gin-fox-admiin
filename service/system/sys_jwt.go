package system

import "gin-one/global"

type JWTService struct{}

func (service *JWTService) IsBlock(token string) bool {
	_, ok := global.BlackCache.Get(token)
	return ok
}

var JWTServiceApp = new(JWTService)
