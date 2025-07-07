package routes

import (
	"booking/rest-api/models"
	"booking/rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user) // Bind the JSON request body to the user struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.Save() // Save the user using the Save method defined in the models package
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user) // Bind the JSON request body to the user struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.ValidatePassword()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID) // Generate a token for the user
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// Here you would typically check the user's credentials against the database
	// For simplicity, we are just returning a success message
	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}
