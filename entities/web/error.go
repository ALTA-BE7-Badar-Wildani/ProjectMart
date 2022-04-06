package web

type WebError struct {
	Code int
	Message string
}

func (err WebError) Error() string {
	return err.Message
}