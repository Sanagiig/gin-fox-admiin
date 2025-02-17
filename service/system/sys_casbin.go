package system

import (
	"gin-one/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"sync"
)

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	onece                sync.Once
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (service *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	onece.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.DB)
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}

		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, err = casbin.NewSyncedCachedEnforcer(m, a)
		if err != nil {
			zap.L().Error("创建Casbin 失败!", zap.Error(err))
			panic(err)
		}
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		err = syncedCachedEnforcer.LoadPolicy()
		if err != nil {
			zap.L().Error("加载Policy失败!", zap.Error(err))
			return
		}
	})

	return syncedCachedEnforcer
}
