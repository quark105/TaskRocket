package config

import (
	"os"

	log "github.com/sebastianpacuk/taskrocket/utils/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Port          string            `mapstructure:"port"`
	TaskSvcUrl    string            `mapstructure:"taskSVCUrl"`
	LoggingConfig log.LoggingConfig `mapstructure:"logging"`
}

func LoadConfig() (c Config, err error) {
	env := os.Getenv("TASKROCKETENV")
	if env == "" {
		env = "dev"
	}

	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")

	viper.SetDefault("port", "9000")
	viper.SetDefault("taskSVCUrl", "localhost:50001")
	viper.SetDefault("logging.service", "TaskRocket-apigateway")
	viper.SetDefault("logging.formatter", "text")
	viper.SetDefault("logging.level", "DEBUG")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}
