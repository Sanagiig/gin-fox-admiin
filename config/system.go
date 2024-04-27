package config

type System struct {
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`          // 端口值
	Lang         string `mapstructure:"lang" json:"lang" yaml:"lang"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}
