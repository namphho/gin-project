package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Login struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title" gorm:"column:title;"`
	Status  int    `json:"status" gorm:"column:status;"`
	Content string `json:"content" gorm:"column:content;"`
}

func main() {

	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	db = db.Debug()

	fmt.Println("open DB success")

	var notes []Note
	db.Where("id = 2").Find(&notes)

	fmt.Println(notes)

	r := gin.Default()
	v1 := r.Group("/v1")
	notesApis := v1.Group("/notes")
	{
		notesApis.GET("", func(context *gin.Context) {
			var notes []Note
			db.Find(&notes)
			context.JSON(http.StatusOK, notes)
		})
		notesApis.POST("/create", func(context *gin.Context) {
			var note Note
			if err := context.ShouldBindJSON(&note); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			db.Create(&note)
			context.JSON(http.StatusOK, note)
		})

	}
	_ = r.Run()
}

type Info struct {
	FirstName string
	LastName  string
	Age       int
}
