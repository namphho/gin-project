package business

import (
	"gin-project/common"
	"gin-project/module/notes/model"
)

type NoteStorage interface {
	ListNote(filter *model.Filter, paging *common.Paging) ([]model.Note, error)
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

func (useCase *ListNoteUseCase) GetAllNotes(filter *model.Filter, paging *common.Paging) ([]model.Note, error) {
	data, err := useCase.store.ListNote(filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("notes", err)
	}
	return data, nil
}
