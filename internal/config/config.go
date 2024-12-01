package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Service
	Storage
	Shards []string
}

type Service struct {
	AllowOrigin string
	APIPrefix   string
	PORT        int
}

type Storage struct {
	Domain          string
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	Region          string
}

func NewConfig(path string) (cfg Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	service := viper.GetStringMap("service")
	storage := viper.GetStringMap("storage")
	postgres := viper.GetStringMap("postgres")

	cfg.Service.AllowOrigin = service["allow.origin"].(string)
	cfg.Service.APIPrefix = service["api.prefix"].(string)
	cfg.Service.PORT = service["port"].(int)

	cfg.Storage.Domain = storage["domain"].(string)
	cfg.Storage.Endpoint = storage["endpoint"].(string)
	cfg.Storage.AccessKeyID = storage["access_key_id"].(string)
	cfg.Storage.SecretAccessKey = storage["secret_access_id"].(string)
	cfg.Storage.Bucket = storage["bucket"].(string)
	cfg.Storage.Region = storage["region"].(string)

	shards := postgres["shards"].([]any)

	cfg.Shards = make([]string, len(shards))
	for i, shard := range shards {
		cfg.Shards[i] = shard.(string)
	}

	return cfg, nil
}
