package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() {
	Router := gin.Default()
	Router.Use(gin.Logger())

	User(Router)

	Router.Run(":3000")
}
