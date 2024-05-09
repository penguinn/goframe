package config

import (
	"flag"

	"github.com/penguinn/go-sdk/config"
)

type appConfig struct {
	DSN        map[string]string `mapstructure:"DSN"`        // MySQL的DSN表达
	ServerPort int               `mapstructure:"ServerPort"` // 启动端口
}

var Config = appConfig{
	ServerPort: 8099,
}

var (
	configPath     string
	configFileType string
)

func Init() error {
	flag.StringVar(&configFileType, "t", "yaml", "config file type")
	flag.StringVar(&configPath, "f", "./config.yaml", "config file path")
	flag.Parse()

	// 优先级，环境变量 > 文件配置 > 默认配置
	err := config.Init(configFileType, configPath, &Config)
	if err != nil {
		return err
	}

	return nil
}
