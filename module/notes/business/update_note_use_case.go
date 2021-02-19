package business

import (
	"gin-project/common"
	"gin-project/module/notes/model"
)

type IUpdateNoteUseCase interface {
	FindNote(id int) (*model.Note, error)
	UpdateNote(noteId int, data *model.NoteUpdate) error
}

type UpdateNoteUseCase struct {
	useCase IUpdateNoteUseCase
}

func NewInstanceUpdateNoteUseCase(useCase IUpdateNoteUseCase) *UpdateNoteUseCase{
	return &UpdateNoteUseCase{useCase: useCase}
}

func (u *UpdateNoteUseCase) UpdateNoteById(noteId int, data *model.NoteUpdate) error{
	iUpdateNoteUseCase := u.useCase
	if _, err := iUpdateNoteUseCase.FindNote(noteId); err != nil {
		return common.ErrCannotGetEntity("can't find note", err)
	}

	if err := iUpdateNoteUseCase.UpdateNote(noteId, data); err != nil {
		return common.ErrCannotUpdateEntity("can't update note", err)
	}
	return nil
}