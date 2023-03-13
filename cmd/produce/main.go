package main

import (
	"fmt"
	"kafka-gaming/api"
	"kafka-gaming/api/config"
	kafkas "kafka-gaming/kafkas"
)

func main() {

	errChan := make(chan error, 1)
	serverConfig := config.NewConfig()

	go func() {
		errChan <- api.Router(serverConfig.ListenAddrApi)
	}()

	go kafkas.RunConsumer(serverConfig)

	err := <-errChan
	if err != nil {
		fmt.Println(err)
	}

}
