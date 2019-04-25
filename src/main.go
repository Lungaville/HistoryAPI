package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/api/user/:id", inDB.GetUser)
	router.GET("/api/users", inDB.GetUsers)
	router.POST("/api/user", inDB.CreateUser)
	router.PUT("/api/user", inDB.UpdateUser)
	router.DELETE("/api/user/:id", inDB.DeleteUser)
	router.Run(":8000")

}
