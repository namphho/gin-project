package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Login struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"user": Login{
				"Hnam", "12345678",
			},
		})
	})

	v1 := r.Group("/v1")
	notes := v1.Group("/notes")
	{
		notes.GET("")
		notes.GET("/:note-id")
		notes.PUT("/:note-id")
		notes.DELETE("/:note-id")
		notes.POST("/login", func(c *gin.Context) {
			var json Login
			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"user": Login{
					json.User + "-resp", json.Password,
				},
			})
		})
	}

	v1.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		age, _ := strconv.Atoi(c.Query("age")) // shortcut for c.Request.URL.Query().Get("lastname")

		//c.String(http.StatusOK, "Hello %s %s %d", firstname, lastname, age)
		c.JSON(http.StatusOK, gin.H{
			"data": Info{
				FirstName: firstname,
				LastName:  lastname,
				Age:       age,
			},
		})
	})

	_ = r.Run()
}

type Info struct {
	FirstName string
	LastName  string
	Age       int
}
