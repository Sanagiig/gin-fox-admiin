package initialize

import (
	"gin-one/global"
	"gin-one/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysRole{},
		system.SysAuthority{},
		system.SysUserRole{},
		system.SysRoleAuthority{},
		system.SysDataDic{},
	)

	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	global.Log.Info("register table success")
}
