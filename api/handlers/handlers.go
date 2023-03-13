package handlers

import (
	"context"
	"fmt"
	"kafka-gaming/api/usecase"
	"kafka-gaming/manager"
	"kafka-gaming/models"

	"github.com/gin-gonic/gin"
)

type handler struct {
	usecase usecase.Usecase
}

type Handlers interface {
	Test(c *gin.Context)
}

func NewHandler() Handlers {
	return &handler{
		usecase: manager.GetUsecase(),
	}
}

func (h *handler) Test(c *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	form := models.RequestTest{}

	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}

	key := "key123"
	fmt.Println([]byte(key))

	err = h.usecase.Test(key, &form)
	if err != nil {
		c.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}

	c.JSON(200, map[string]string{
		"msg": "success",
	})

}
