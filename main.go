package main

import (
	"chat-app/handler"
	"chat-app/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){

  dsn := "laravel:password@tcp(127.0.0.1:3306)/chat?charset=utf8&parseTime=True"
  db, err := gorm.Open(mysql.Open(dsn) , &gorm.Config{})
  if err != nil {
    log.Fatal(err.Error())
  }
  userRepository := user.NewRepository(db)
  userService := user.NewService(userRepository)
  userHandler := handler.NewUserHandler(userService)

  router := gin.Default()
  api:= router.Group("/api/v1")
  api.POST("/users",userHandler.RegisterUser)
  router.Run()
}