package business

import (
	"gin-project/module/common"
	"gin-project/module/notes/model"
)

type ICreateNoteUseCase interface {
	Create(content *model.NoteCreate) error
}

type CreateNoteUseCase struct {
	useCase ICreateNoteUseCase
}

func NewInstanceCreateNoteUseCase(useCase ICreateNoteUseCase) *CreateNoteUseCase{
	return &CreateNoteUseCase{useCase: useCase}
}

func (c *CreateNoteUseCase) CreateNote(data *model.NoteCreate) error {
	err := c.useCase.Create(data)
	if err != nil {
		return common.ErrCannotCreateEntity(model.Note{}.EntityName(), err)
	}
	return nil
}