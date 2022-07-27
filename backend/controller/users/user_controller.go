package users

import (
	"github.com/henry-on-the-internet/go-react-auth-demo/backend/domain/users"
	"github/com/gin-gonic/gin"
)
func Register(c *gin.Context){
	var users users.User


	err := c.ShouldBindJSON(&user)
	if err != nil {
		errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return 
	}

}