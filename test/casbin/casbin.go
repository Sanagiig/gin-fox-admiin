package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"os"
	"path"
	"sync"
)

var syncedCachedEnforcer *casbin.Enforcer
var once sync.Once

func GetCasbin() *casbin.Enforcer {
	once.Do(func() {
		var err error
		pwd, _ := os.Getwd()

		m := path.Join(pwd, "/test/casbin/model.config")
		p := path.Join(pwd, "/test/casbin/policy.csv")
		syncedCachedEnforcer, err = casbin.NewEnforcer(m, p)
		if err != nil {
			panic(err)
		}

		err = syncedCachedEnforcer.LoadPolicy()
		if err != nil {
			zap.L().Error("加载Policy失败!", zap.Error(err))
			return
		}
	})
	return syncedCachedEnforcer
}

func main() {
	cas := GetCasbin()
	str := cas.GetAllActions()
	cas.AddPolicy("LWJ", "book", "read")

	ok, err := cas.AddRolesForUser("LWJ", []string{"ADMIN", "TEST"})
	fmt.Println(ok, err)
	if err != nil {
		panic(err)
	}

	res, _ := cas.GetRolesForUser("LWJ")
	fmt.Println("role for user", res)

	res, _ = cas.GetDomainsForUser("LWJ")
	fmt.Println("doemal for user", res)

	err = cas.SavePolicy()
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
