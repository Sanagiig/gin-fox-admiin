package global

import (
	"gin-one/config"
	"gin-one/message"
	"github.com/eko/gocache/store"
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	DbMap      map[string]*gorm.DB
	Config     config.Config
	Log        *zap.Logger
	Viper      *viper.Viper
	Mongo      *qmgo.QmgoClient
	Msg        *message.Message
	BlackCache store.GoCacheClientInterface
)
