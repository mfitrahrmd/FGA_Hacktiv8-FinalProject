package photo

type PHOTO_ERROR string

const (
	PHOTO_NOT_FOUND PHOTO_ERROR = "photo not found"
	ACCESS_DENIED   PHOTO_ERROR = "access denied"
)

type photoError struct {
	message PHOTO_ERROR
}

func (u photoError) Error() string {
	return string(u.message)
}

func NewPhotoError(message PHOTO_ERROR) photoError {
	return photoError{
		message: message,
	}
}
