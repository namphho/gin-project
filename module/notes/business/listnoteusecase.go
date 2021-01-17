package business

import (
	"gin-project/module/notes/model"
)

type NoteStorage interface {
	ListNote() ([]model.Note, error)
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

func (useCase *ListNoteUseCase) GetAllNotes() ([]model.Note, error) {
	return useCase.store.ListNote()
}
