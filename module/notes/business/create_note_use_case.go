package business

import (
	"gin-project/common"
	"gin-project/module/notes/model"
)

type ICreateNoteUseCase interface {
	Create(content *model.NoteCreate) error
}

type CreateNoteUseCase struct {
	useCase ICreateNoteUseCase
	requester common.Requester
}

func NewInstanceCreateNoteUseCase(useCase ICreateNoteUseCase, requester common.Requester) *CreateNoteUseCase{
	return &CreateNoteUseCase{useCase: useCase, requester: requester}
}

func (c *CreateNoteUseCase) CreateNote(data *model.NoteCreate) error {
	data.UserId = c.requester.GetUserId()
	err := c.useCase.Create(data)
	if err != nil {
		return common.ErrCannotCreateEntity(model.Note{}.EntityName(), err)
	}
	return nil
}