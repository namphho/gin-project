package business

import (
	"gin-project/module/common"
	"gin-project/module/notes/model"
)

type NoteStorage interface {
	ListNote(paging *common.Paging) ([]model.Note, error)
}

type NoteStorageFake interface {
	ListNote() ([]model.Note, error)
	DeleteNote()
}

type ListNoteUseCase struct {
	store NoteStorage
}

func NewInstance(storage NoteStorage) *ListNoteUseCase {
	return &ListNoteUseCase{storage}
}

func (useCase *ListNoteUseCase) GetAllNotes(paging *common.Paging) ([]model.Note, error) {
	return useCase.store.ListNote(paging)
}
