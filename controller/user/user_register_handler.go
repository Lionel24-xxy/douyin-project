package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"TikTok_Project/service/user"
)

func UserRegister(c *gin.Context){
	name := c.Query("username")
	password := c.Query("password")
	
	registerResponse, statusCode, err := user.UserRegister(c, name, password)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": statusCode, 
			"status_msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode, 
		"status_msg": "Register succeed!",
		"user_id": registerResponse.UserId,
		"token": registerResponse.Token,
	})
}

