package main

import (
	"fmt"
	"gin-project/module/appctx"
	"gin-project/module/notes/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Login struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type FakeStore struct{}

func (FakeStore) ListNote() ([]model.Note, error) {
	return []model.Note{
		model.Note{
			Title:   "title test",
			Content: "content test",
		},
	}, nil
}

func main() {

	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	db = db.Debug()
	appCtx := appctx.NewInstance(db)

	fmt.Println("open DB success")

	r := gin.Default()
	setUpRouter(r, appCtx)

	_ = r.Run()
}

type Info struct {
	FirstName string
	LastName  string
	Age       int
}
