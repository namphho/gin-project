package business

import (
	"errors"
	"gin-project/module/notes/model"
)

type IGetNoteUseCase interface {
	FindNoteWithKeys(id int, moreKeys ...string) (*model.Note, error)
}

type GetNoteUseCase struct {
	usecase IGetNoteUseCase
}

func NewInstanceGetNoteUseCase(useCase IGetNoteUseCase) *GetNoteUseCase {
	return &GetNoteUseCase{useCase}
}

func (impl *GetNoteUseCase) GetNote(noteId int) (*model.Note, error) {
	//find note
	note, err := impl.usecase.FindNoteWithKeys(noteId, "User")

	if err != nil {
		return nil, errors.New("note not found")
	}

	return note, nil
}
