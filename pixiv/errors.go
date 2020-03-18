package pixiv

import "fmt"

type APIError struct {
	Message string
	Code    int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("pixiv: %d %s", e.Code, e.Message)
}
