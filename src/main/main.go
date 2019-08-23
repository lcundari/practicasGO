package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcundarigit/practicasGO/src/service"
)

var miTweet *service.TweetManager

func main() {
	router := gin.Default()
	miTweet = service.NewTweetManager()
	router.GET("/GET", funcionQueHaceGet)
	// router.POST("/POST", funcionQueHacePost)

	router.Run("localhost:8080")
}
func funcionQueHaceGet(c *gin.Context) {
	c.JSON(200, miTweet.GetTweets())

}

// func funcionQueHacePost(c *gin.Context) {
// 	parametro := c.Param("parametro")
// 	c.String(http.StatusOK, "ok")
// }
