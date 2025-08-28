
package conf

import (
    "github.com/go-keg/keg/contrib/config"
)

type Crm struct {
    Key    string
    Server struct {
        HTTP config.Server `yaml:"http"`
    }
    Data struct {
        Database           config.Database
        Kafka              config.Kafka
        KafkaConsumerGroup config.KafkaConsumerGroup
    }
    Trace struct {
        Endpoint string
    }
    Log config.Log
}

func Load(path string, envs ...string) (*Crm, error) {
    return config.Load[Crm](path, envs...)
}

func MustLoad(path string, envs ...string) *Crm {
    cfg, err := Load(path, envs...)
    if err != nil {
        panic(err)
    }
    return cfg
}
