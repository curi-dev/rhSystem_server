package error

type AppError struct {
	Err        error
	Message    string
	StatusCode int
}
