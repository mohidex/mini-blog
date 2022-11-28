package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/mini-blog/helper"
	"github.com/mohidex/mini-blog/model"
	"gorm.io/gorm"
)

func AddBlog(ctx *gin.Context) {
	var input model.Blog
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedBlog, err := input.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": savedBlog})
}

func GetAllBlog(ctx *gin.Context) {
	user, err := helper.CurrentUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user.Blogs})
}

func GetBlogById(ctx *gin.Context) {
	id := ctx.Param("id")
	blog, err := model.FindBlogById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": blog})
}

func DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	blog, err := model.FindBlogById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	if err := blog.Delete(); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"status": "ok"})
}
