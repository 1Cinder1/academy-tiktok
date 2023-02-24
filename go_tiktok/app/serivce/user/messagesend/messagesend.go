package messagesend

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/messagesend/internal"
	"go_tiktok/utils/jwt"
	"net/http"
	"strconv"
)

func SendMessage(c *gin.Context) {

	token := c.Query("token")
	userId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	content := c.Query("content")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  "user_id cannot be null",
		})
		return
	}
	if actionType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  "actionType cannot be null",
		})
		return
	}

	id, err := jwt.GetUserId(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return

	}
	userid, _ := strconv.ParseInt(userId, 10, 64)
	err1 := internal.SendMessage(c, id, int(userid), content)
	if err1 != nil {

		if err1.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": -1,
				"status_msg":  err1.Error(),
			})
			return
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  http.StatusOK,
	})
	return

}
