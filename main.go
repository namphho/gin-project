package main

import (
	"fmt"
	"gin-project/module/appctx"
	"gin-project/module/notes/business"
	"gin-project/module/notes/model"
	"gin-project/module/notes/storage"
	"gin-project/module/notes/transport"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
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
	v1 := r.Group("/v1")
	notesApis := v1.Group("/notes")
	{
		notesApis.GET("", transport.GetNotes(appCtx))
		notesApis.POST("/create", transport.CreateNote(appCtx))

		notesApis.DELETE("/:note-id", func(ctx *gin.Context) {
			idString := ctx.Param("note-id")
			id, _ := strconv.Atoi(idString)

			mysqlStorage := storage.NewInstance(db)
			useCase := business.NewInstanceDeleteUseCase(mysqlStorage)
			if err := useCase.DeleteNote(id); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"data": "okay"})
		})

	}
	_ = r.Run()
}

type Info struct {
	FirstName string
	LastName  string
	Age       int
}
