package config

import (
	"os"

	log "github.com/sebastianpacuk/taskrocket/utils/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Port          string            `mapstructure:"port"`
	DBUrl         string            `mapstructure:"dbUrl"`
	LoggingConfig log.LoggingConfig `mapstructure:"logging"`
}

func LoadConfig() (config Config, err error) {
	env := os.Getenv("TASKROCKETENV")
	if env == "" {
		env = "dev"
	}

	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")

	viper.SetDefault("port", "50001")
	viper.SetDefault("dbUrl", "postgres://postgres:postgres@localhost:5432/tasks")
	viper.SetDefault("logging.service", "TaskRocket-tasks")
	viper.SetDefault("logging.formatter", "text")
	viper.SetDefault("logging.level", "INFO")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
