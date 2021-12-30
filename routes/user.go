package routes

import (
	"math/rand"
	"net/http"
	"server/db"
	"server/model"

	"github.com/gin-gonic/gin"
)

func User(DefaultRouter *gin.Engine) {
	router := DefaultRouter.Group("/user")

	router.GET("/", func(c *gin.Context) {
		db.Database.Create(&model.User{
			Name: "jinzhu",
			Age:  rand.Uint32(),
		})
		Users := &[]model.User{}

		db.Database.Where("name = ?", "jinzhu").Find(Users)

		c.JSON(http.StatusOK, Users)
	})
}
