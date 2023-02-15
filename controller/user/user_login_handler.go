package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"TikTok_Project/service/user"
)


func UserLogin(c *gin.Context){
	name := c.Query("username")
	password := c.Query("password")
	
	loginResponse, statusCode, err := user.UserLogin(name, password)
	
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": statusCode,
			"status_msg": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode, 
		"status_msg": "Login succeed!",
		"user_id": (*loginResponse).UserId,
		"token": (*loginResponse).Token,
	})
}