package service

import "errors"

type ServiceError struct {
	StatusCode int
	Err        error
}

var (
	NOT_FOUND                       ServiceError = ServiceError{404, errors.New("not found")}
	ACCESS_DENIED                   ServiceError = ServiceError{403, errors.New("access denied")}
	INVALID_PASSWORD                ServiceError = ServiceError{401, errors.New("invalid password")}
	USERNAME_OR_EMAIL_ALREADY_EXIST ServiceError = ServiceError{400, errors.New("username or email already exist")}
	USER_DOES_NOT_EXIST             ServiceError = ServiceError{404, errors.New("user does not exist")}
	PHOTO_NOT_FOUND                 ServiceError = ServiceError{404, errors.New("photo not found")}
	COMMENT_NOT_FOUND               ServiceError = ServiceError{404, errors.New("comment not found")}
	SOCIAL_MEDIA_NOT_FOUND          ServiceError = ServiceError{404, errors.New("social media not found")}
)

func (s ServiceError) Error() string {
	return s.Err.Error()
}
