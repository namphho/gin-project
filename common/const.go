package common

const (
	CurrentUser = "current_user"
)

type Masker interface {
	Mask(isAdmin bool)
}

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
