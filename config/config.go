package config

type Config struct {
	MySql MySql `mapstructure:"mysql`
}

type MySql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url string `mapstructure:"url"`
}