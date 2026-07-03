package config

import (
	"os"
	"strconv"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Application ApplicationConfig `yaml:"application"`
	Server      ServerConfig      `yaml:"server"`
	MongoDB     MongoDBConfig     `yaml:"mongodb"`
	Kafka       KafkaConfig       `yaml:"kafka"`
}

type ApplicationConfig struct {
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type MongoDBConfig struct {
	URI        string `yaml:"uri"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

type KafkaConfig struct {
	Broker string `yaml:"broker"`
	Topic  string `yaml:"topic"`
}

var (
	cfg  *Config
	once sync.Once
)

func Load(path string) (*Config, error) {

	var err error

	once.Do(func() {

		file, e := os.ReadFile(path)

		if e != nil {
			err = e
			return
		}

		config := &Config{}

		e = yaml.Unmarshal(file, config)

		if e != nil {
			err = e
			return
		}

		// ----------------------------
		// Environment Variable Override
		// ----------------------------

		if value := os.Getenv("APP_NAME"); value != "" {
			config.Application.Name = value
		}

		if value := os.Getenv("APP_ENV"); value != "" {
			config.Application.Environment = value
		}

		if value := os.Getenv("SERVER_PORT"); value != "" {

			port, e := strconv.Atoi(value)

			if e == nil {
				config.Server.Port = port
			}
		}

		if value := os.Getenv("MONGO_URI"); value != "" {
			config.MongoDB.URI = value
		}

		if value := os.Getenv("MONGO_DATABASE"); value != "" {
			config.MongoDB.Database = value
		}

		if value := os.Getenv("MONGO_COLLECTION"); value != "" {
			config.MongoDB.Collection = value
		}

		if value := os.Getenv("KAFKA_BROKER"); value != "" {
			config.Kafka.Broker = value
		}

		if value := os.Getenv("KAFKA_TOPIC"); value != "" {
			config.Kafka.Topic = value
		}

		cfg = config
	})

	return cfg, err
}

func Get() *Config {
	return cfg
}
