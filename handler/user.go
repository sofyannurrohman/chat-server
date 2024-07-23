package handler

import (
	"chat-app/helper"
	"chat-app/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func(h *userHandler) RegisterUser(c *gin.Context){
	//get request user
	var input user.RegisterUserInput

	err:=c.ShouldBindJSON(&input)
	if err != nil {
		var errors = helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}
		response:= helper.APIResponse("Register account failed",http.StatusUnprocessableEntity,"error",errorMessage)
		 c.JSON(http.StatusBadRequest,response)
		 return
	}
	newUser,err:= h.userService.RegisterUser(input)
	
	if err != nil {

		response:= helper.APIResponse("Register account failed",http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest,response)	
		return
	}
	response:= helper.APIResponse("Account has been registered",http.StatusOK,"success",newUser)
	c.JSON(http.StatusOK,response)

	
	//map input form user to RegisterInputUser
	//return will using as param to service
}