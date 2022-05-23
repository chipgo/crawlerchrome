package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Service      Service       `mapstructure:"service"`
	CronSchedule CronSchedules `mapstructure:"cron_schedules"`
}

type (
	Service struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
	}

	CronSchedules struct {
		Crawler string `mapstructure:"crawler"`
	}
)

func NewConfig() *Config {

	stage := "local"
	if s := os.Getenv("APP_STAGE"); len(s) != 0 {
		stage = s
	}
	filename := fmt.Sprintf("cfg.%s", stage)

	log.Printf("config filename: %s\n", filename)
	viper.AddConfigPath("./resources/") // path to look for the config file in
	viper.SetConfigName(filename)       // name of config file (without extension)
	viper.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")            // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panic(errors.New("config not found"))
		} else {
			log.Panic(err)
		}
	}

	cfg := new(Config)
	if err := viper.Unmarshal(cfg); err != nil {
		log.Panic(err)
	}

	log.Printf("config value: %+v\n", *cfg)
	return cfg
}
