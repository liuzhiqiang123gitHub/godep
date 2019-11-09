package live_class

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
func main()  {
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message" :"maruiyi",
		})
		fmt.Println("hello,world")

	})

	router.Run(":8080")
}
