package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/mini-blog/model"
)

func Register(ctx *gin.Context) {
	var input model.RegistrationInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	user := model.User{
		Name: input.Name,
		Username: input.Username,
		Email: input.Email,
		Password: input.Password,
	}
	savedUser, err := user.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user": savedUser,
	})
}
