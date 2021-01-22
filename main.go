package main

import (
	"fmt"
	"gin-project/module/notes/business"
	"gin-project/module/notes/model"
	"gin-project/module/notes/storage"
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

	fmt.Println("open DB success")

	r := gin.Default()
	v1 := r.Group("/v1")
	notesApis := v1.Group("/notes")
	{
		notesApis.GET("", func(context *gin.Context) {
			mysqlStorage := storage.NewInstance(db)
			listNoteUseCase := business.NewInstance(mysqlStorage)
			notes, _ := listNoteUseCase.GetAllNotes()

			context.JSON(http.StatusOK, notes)
		})
		notesApis.POST("/create", func(context *gin.Context) {
			var note model.Note
			if err := context.ShouldBindJSON(&note); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			db.Create(&note)
			context.JSON(http.StatusOK, note)
		})

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
