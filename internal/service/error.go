package service

type SERVICE_ERROR string

const (
	NOT_FOUND        SERVICE_ERROR = "not found"
	ACCESS_DENIED    SERVICE_ERROR = "access denied"
	INVALID_PASSWORD SERVICE_ERROR = "invalid password"
)

type serviceError struct {
	message SERVICE_ERROR
}

func (s serviceError) Error() string {
	return string(s.message)
}

func NewServiceError(message SERVICE_ERROR) serviceError {
	return serviceError{
		message: message,
	}
}
