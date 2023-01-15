package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SingUp(context *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if context.Bind(&err) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed hash password",
		})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func Login(context *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed read body",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	context.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func Validate(context *gin.Context) {
	user, _ := context.Get("user")

	context.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
