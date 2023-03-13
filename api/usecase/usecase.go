package usecase

import (
	"fmt"
	"kafka-gaming/api/config"
	"kafka-gaming/kafkas"

	managerusecase "kafka-gaming/managerUsecase"
	"kafka-gaming/models"

	"gopkg.in/square/go-jose.v2/json"
)

type usecase struct {
	config config.ServerConfig
}

type Usecase interface {
	Test(key string, form *models.RequestTest) error
}

func NewUsecase() Usecase {
	return &usecase{
		config: managerusecase.GetConfig(),
	}
}

func (uc *usecase) Test(key string, form *models.RequestTest) error {
	fmt.Println(uc.config)

	val, _ := json.Marshal(form.Text)
	err := kafkas.Produce(uc.config, key, string(val))
	if err != nil {
		return err
	}

	return nil
}
