package initialize

import (
	"gin-one/global"
	"github.com/patrickmn/go-cache"
	"time"
)

func InitializeGlobal() {
	global.DB = Gorm()
	global.Msg = Msg()
	global.BlackCache = cache.New(5*time.Minute, 10*time.Minute)
	DBList()
}
