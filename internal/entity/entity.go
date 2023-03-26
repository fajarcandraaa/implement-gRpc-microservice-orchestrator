package entity

type Error string

// Declare error messege
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")

	//User Error
	ErrUserNotExist            = Error("domain.user.error.not_exist")
	ErrUserAlreadyExist        = Error("domain.user.error.email_or_username_alredy_exist")
	ErrUsersCredentialNotExist = Error("domain.user.error.credential_not_exist")

	//Book Error
	ErrBookNotExist            = Error("domain.book.error.not_exist")
	ErrBookAlreadyExist        = Error("domain.book.error.book_title_already_exist")
	ErrBooksCredentialNotExist = Error("domain.book.error.credential_not_exist")
)

func (e Error) Error() string {
	return string(e)
}
