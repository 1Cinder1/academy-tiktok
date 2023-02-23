package actionList

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/actionlist/internal"
	"go_tiktok/utils/jwt"
	"net/http"
	"strconv"
)

func GetVideoList(c *gin.Context) {

	token := c.Query("token")
	userId := c.Query("user_id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  "token or user_id cannot be null",
			"video_list":  "null",
		})
		return
	}
	if token != "" {
		id, err := jwt.GetUserId(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": -1,
				"status_msg":  err.Error(),
				"video_list":  "null",
			})
			return

		}
		userid, _ := strconv.ParseInt(userId, 10, 64)
		err, videolist := internal.GetVideoList(c, int(userid), id)
		if err != nil {

			if err.Error() == "internal err" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": -1,
					"status_msg":  err.Error(),
					"video_list":  "null",
				})
				return
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  http.StatusOK,
			"video_list":  videolist,
		})
		return
	} else {
		id := -1
		userid, _ := strconv.ParseInt(userId, 10, 64)
		err, viseolist := internal.GetVideoList(c, int(userid), id)
		if err != nil {

			if err.Error() == "internal err" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": -1,
					"status_msg":  err.Error(),
					"video_list":  "null",
				})
				return
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  http.StatusOK,
			"video_list":  viseolist,
		})
		return
	}

}
