package usercontainer

import (
	usermodel "admodev/invexchange/model"
	timetool "admodev/invexchange/tool"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserInput struct {
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	DocumentNumber string `json:"document_number" binding:"required"`
	Birthdate      string `json:"birthdate" binding:"required"`
}

func GetUsers() {
	fmt.Println("todo...")
}

func storeUserInDatabase(user usermodel.User) {
	fmt.Println("todo...")
}

func CreateUser(c *gin.Context) {
	var userInput createUserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	user := usermodel.User{FirstName: userInput.FirstName, LastName: userInput.LastName, DocumentNumber: userInput.DocumentNumber, Birthdate: timetool.ParseDate(userInput.Birthdate)}

	log.Println("INFO: User created.")

	storeUserInDatabase(user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}
