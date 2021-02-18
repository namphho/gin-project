package business

import (
	"errors"
	"gin-project/module/notes/model"
)

type IGetNoteUseCase interface {
	FindNote(id int) (*model.Note, error)
}

type GetNoteUseCase struct {
	usecase IGetNoteUseCase
}

func NewInstanceGetNoteUseCase(useCase IGetNoteUseCase) *GetNoteUseCase {
	return &GetNoteUseCase{useCase}
}

func (impl *GetNoteUseCase) GetNote(noteId int) (*model.Note, error) {
	//find note
	note, err := impl.usecase.FindNote(noteId)

	if err != nil {
		return nil, errors.New("note not found")
	}

	return note, nil
}
