package user

type USER_ERROR string

const (
	EMAIL_OR_USERNAME_ALREADY_EXIST USER_ERROR = "email or username already exist"
	USER_NOT_FOUND                  USER_ERROR = "user not found"
	INVALID_PASSWORD                USER_ERROR = "invalid password"
)

type userError struct {
	message USER_ERROR
}

func (u userError) Error() string {
	return string(u.message)
}

func NewUserError(message USER_ERROR) userError {
	return userError{
		message: message,
	}
}
