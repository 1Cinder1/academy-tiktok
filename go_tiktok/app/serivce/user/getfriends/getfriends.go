package getfriends

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/getfriends/internal"

	"net/http"
	"strconv"
)

func GetFriends(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	id := 47
	if token == "" || userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  "token or user_id cannot be null",
			"user":        "null",
		})
		return
	}

	userid, _ := strconv.ParseInt(userId, 10, 64)
	err, user := internal.GetFriends(c, int(userid), id)
	if err != nil {

		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": -1,
				"status_msg":  err.Error(),
				"user":        "null",
			})
			return
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
