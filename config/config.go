package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
