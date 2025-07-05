package routes

import (
	"booking/rest-api/models"
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
