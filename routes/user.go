package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/akshrana12/Library-Management/controller"
	models "github.com/akshrana12/Library-Management/model"
	"github.com/gin-gonic/gin"
)

func guidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var obj models.Users
		if errA := c.BindJSON(&obj); errA == nil {
			if strings.Contains(obj.EmailId, "@geu.ac.in") || strings.Contains(obj.EmailId, "@gmail.com") {
				b, _ := json.Marshal(obj)
				c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong email Handle"})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}

	}
}
func UserRoutes(router *gin.Engine) {
	bookroute := router.Group("api/user")
	bookroute.POST("/addUser", guidMiddleware(), controller.AddUser)
	bookroute.GET("/getAllUsers", controller.ShowUsers)
	bookroute.GET("/getUser/:id", controller.GetUserById)
}
