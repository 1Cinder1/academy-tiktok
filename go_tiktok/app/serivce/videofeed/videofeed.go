package videofeed

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VideoFeed(context *gin.Context) {
	//ginServer := gin.Default()
	//ginServer.GET("/douyin/feed", myHandler(), func(context *gin.Context) {
	latest_time := context.Query("update_time")
	token := context.Query("token")
	create_time := context.Query("create_time")

	video_list := VideoFeed

	if latest_time != "" && token == "" {
		context.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "not logged in",
			"next_time":   create_time,
			"video_list":  video_list,
		})
		return
	} else if latest_time != "" && token != "" {
		context.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "have logged in",
			"next_time":   create_time,
			"video_list":  video_list,
		})
		return
	}

	//})
}
