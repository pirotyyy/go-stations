package model

type ErrNotFound struct {
	Message string
}

func (e *ErrNotFound) Error() string {
	return "not found"
}

type ErrUnauthorized struct {
	Message string `json:"message"`
}
