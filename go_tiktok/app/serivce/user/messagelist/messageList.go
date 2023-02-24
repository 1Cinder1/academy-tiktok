package messagelist

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/messagelist/internal"
	"go_tiktok/utils/jwt"
	"net/http"
	"strconv"
)

func GetMessageList(c *gin.Context) {

	token := c.Query("token")
	userId := c.Query("to_user_id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code":  -1,
			"status_msg":   "token or user_id cannot be null",
			"message_list": "null",
		})
		return
	}

	id, err := jwt.GetUserId(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code":  -1,
			"status_msg":   err.Error(),
			"message_list": "null",
		})
		return

	}
	userid, _ := strconv.ParseInt(userId, 10, 64)
	err, messagelist := internal.GetMessageList(c, int(userid), id)
	if err != nil {

		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code":  -1,
				"status_msg":   err.Error(),
				"message_list": "null",
			})
			return
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   http.StatusOK,
		"message_list": messagelist,
	})
	return

}
