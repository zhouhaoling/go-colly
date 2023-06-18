package initialize

import (
	"colly_music/global"
	"fmt"

	"github.com/spf13/viper"
)

//读取配置文件
func LoadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error resources file: %v\n", err))
	}

	err1 := viper.Unmarshal(&global.Config)
	if err1 != nil {
		panic(fmt.Errorf("unable to decode into struct %v\n", err1))
	}
}
