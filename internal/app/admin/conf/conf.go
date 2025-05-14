package conf

import (
	"github.com/go-keg/keg/contrib/config"
)

type Config struct {
	Key    string
	Name   string
	Server struct {
		Http config.Server
	}
	Data struct {
		Database config.Database
		Kafka    config.Kafka
	}
	OAuth struct {
		Google struct {
			RedirectURL 	string
			ClientID 		string
			ClientSecret 	string
		}
	}
	KafkaConsumerGroup config.KafkaConsumerGroup
	Email              config.Email
	Trace              struct {
		Endpoint string
	}
	Log config.Log
}

func Load(path string, envs ...string) (*Config, error) {
	return config.Load[Config](path, envs...)
}

func MustLoad(path string, envs ...string) *Config {
	cfg, err := Load(path, envs...)
	if err != nil {
		panic(err)
	}
	return cfg
}
