package config

import "github.com/spf13/viper"

func InitConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("未找到配置文件")
		} else {
			panic("加载配置文件失败")
		}
	}
	// 监听配置文件变化
	viper.WatchConfig()
}
