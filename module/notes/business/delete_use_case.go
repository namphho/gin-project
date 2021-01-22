package business

import (
	"errors"
	"gin-project/module/notes/model"
)

type DeleteUseCase interface {
	FindNote(id int) (*model.Note, error)
	DeleteNote(id int) error
}

type DeleteUseCaseImpl struct {
	usecase DeleteUseCase
}

func NewInstanceDeleteUseCase(useCase DeleteUseCase) *DeleteUseCaseImpl {
	return &DeleteUseCaseImpl{useCase}
}

func (impl *DeleteUseCaseImpl) DeleteNote(noteId int) error {
	//find note
	note, err := impl.usecase.FindNote(noteId)

	if err != nil {
		return errors.New("note not found")
	}

	if note.Status == 0 {
		return errors.New("note is deleted")
	}
	//end delete
	if err := impl.usecase.DeleteNote(noteId); err != nil {
		return errors.New("delete error is deleted")
	}
	return nil
}
