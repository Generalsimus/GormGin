package main

import (
	"server/db"
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	db.ConnectDb()
	routes.CreateRouter()

}
