package main

import (
	"fmt"
	"gin-project/appctx"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	dsn := os.Getenv("DB_CONNECTION_STRING")
	secret := os.Getenv("SECRET_KEY")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	db = db.Debug()
	appCtx := appctx.NewInstance(db, secret)

	fmt.Println("open DB success")

	r := gin.Default()
	setUpRouter(r, appCtx)

	_ = r.Run()
}