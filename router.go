package main

import (
	"gin-project/appctx"
	"gin-project/middleware"
	"gin-project/module/notes/transport"
	"github.com/gin-gonic/gin"
)

func setUpRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	notesApis := v1.Group("/notes")
	{
		notesApis.GET("", transport.GetNotes(appCtx))
		notesApis.GET("/:note-id", transport.GetNoteById(appCtx))
		notesApis.POST("/create", transport.CreateNote(appCtx))
		notesApis.DELETE("/:note-id", transport.DeleteNote(appCtx))
		notesApis.PUT("/:note-id", transport.UpdateNote(appCtx))
	}
}
