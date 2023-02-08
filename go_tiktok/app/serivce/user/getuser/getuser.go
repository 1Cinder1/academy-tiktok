package getuser

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/getuser/internal"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	id := 11
	if token == "" || userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  "token or user_id cannot be null",
			"user":        "null",
		})
		return
	}

	userid, _ := strconv.ParseInt(userId, 10, 64)
	err, user := internal.Getuser(c, int(userid), id)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": -1,
				"status_msg":  err.Error(),
				"user":        "null",
			})
		case "user doesn't exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": -1,
				"status_msg":  err.Error(),
				"user":        "null",
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  http.StatusOK,
		"user":        user,
	})
	return
}
