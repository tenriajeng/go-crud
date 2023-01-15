package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostIndex(context *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	context.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func PostCreate(context *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	context.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		context.Status(400)
	}

	context.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func PostShow(context *gin.Context) {
	id := context.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	context.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func PostUpdate(context *gin.Context) {
	id := context.Param("id")

	var body struct {
		Title string
		Body  string
	}

	context.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	context.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func PostDelete(context *gin.Context) {
	id := context.Param("id")
	initializers.DB.Delete(&models.Post{}, id)

	context.JSON(http.StatusOK, gin.H{
		"data": "success delete",
	})
}
