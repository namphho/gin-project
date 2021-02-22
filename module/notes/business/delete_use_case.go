package business

import (
	"errors"
	"gin-project/common"
	"gin-project/module/notes/model"
)

type DeleteUseCase interface {
	FindNote(id int) (*model.Note, error)
	DeleteNote(id int) error
}

type DeleteUseCaseImpl struct {
	usecase   DeleteUseCase
	requester common.Requester
}

func NewInstanceDeleteUseCase(useCase DeleteUseCase, requester common.Requester) *DeleteUseCaseImpl {
	return &DeleteUseCaseImpl{useCase, requester}
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

	//check permission
	isAuthor := impl.requester.GetUserId() == note.UserId
	isAdmin := impl.requester.GetRole() == common.RoleAdmin

	if !isAuthor && !isAdmin {
		return common.ErrNoPermission(nil)
	}

	//end delete
	if err := impl.usecase.DeleteNote(noteId); err != nil {
		return errors.New("delete error is deleted")
	}
	return nil
}
