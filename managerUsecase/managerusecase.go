package managerusecase

import (
	"kafka-gaming/api/config"

	"sync"
)

var oneUc sync.Once
var conf config.ServerConfig

func GetConfig() config.ServerConfig {
	oneUc.Do(func() {
		conf = *config.NewConfig()
	})

	return conf
}
