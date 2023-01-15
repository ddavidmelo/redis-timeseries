package config

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

const cfgFile = "config/config.toml"

var c Config

type General struct {
	Version                      string        `mapstructure:"version"`
	Log_level                    uint8         `mapstructure:"log_level"`
	Log_colors                   bool          `mapstructure:"log_disable_colors"`
	Log_timestamp                bool          `mapstructure:"log_disable_timestamp"`
	PublishRate                  time.Duration `mapstructure:"pub_rate"`
	RetentionSampleDuration      time.Duration `mapstructure:"retention_sample_duration"`
	RetentionAggregationDuration time.Duration `mapstructure:"retention_aggregation_duration"`
	AggregationBucketDuration    time.Duration `mapstructure:"aggregation_bucket_duration"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type Config struct {
	General General
	Redis   Redis
}

func init() {
	viper.Set("general.version", "v.0.1.1")
	viper.SetDefault("general.log_level", 2)
	viper.SetDefault("general.log_disable_colors", false)
	viper.SetDefault("general.log_disable_timestamp", false)
	viper.SetDefault("general.pub_rate", "100ms")
	viper.SetDefault("general.retention_sample_duration", "1m")
	viper.SetDefault("general.retention_aggregation_duration", "5m")
	viper.SetDefault("general.aggregation_bucket_duration", "1m")

	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.password", "")

	init_config()
	viper.WriteConfigAs(cfgFile)
}

func init_config() {
	b, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Error("error loading config file")
	}
	viper.SetConfigType("toml")
	if err := viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
		log.Error("error loading config file")
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Error("unmarshal config file error")
	}

	log.SetFormatter(&log.TextFormatter{
		DisableColors:    c.General.Log_colors,
		DisableTimestamp: c.General.Log_timestamp,
	})

	if c.General.Log_level == 3 {
		log.SetLevel(logrus.DebugLevel)
	} else if c.General.Log_level == 2 {
		log.SetLevel(logrus.InfoLevel)
	} else if c.General.Log_level == 1 {
		log.SetLevel(logrus.WarnLevel)
	} else {
		log.SetLevel(logrus.ErrorLevel)
	}
	log.Debug("config object: ", c)
}

func Init_config() {
	log.Info("config version ", c.General.Version)
}

func GetRedisConfig() Redis {
	return c.Redis
}

func GetGeneralConfig() General {
	return c.General
}
