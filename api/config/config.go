package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	ListenAddrApi string

	// kafka
	KafkaBrokers []string
	KafkaTopic   string
}

func NewConfig() *ServerConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	brokers := strings.Split(os.Getenv("KafkaBroker"), ",")

	return &ServerConfig{
		ListenAddrApi: os.Getenv("ListenAddrApi"),
		KafkaBrokers:  brokers,
		KafkaTopic:    os.Getenv("KafkaTopic"),
	}
}
