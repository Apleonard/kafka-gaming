package manager

import (
	"kafka-gaming/api/usecase"
	"sync"
)

var oneUc sync.Once
var uc usecase.Usecase

func GetUsecase() usecase.Usecase {
	oneUc.Do(func() {
		uc = usecase.NewUsecase()
	})

	return uc
}
