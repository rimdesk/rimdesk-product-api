package common

type ApiResponse[T any] struct {
	Success   bool     `json:"success"`
	Message   string   `json:"message"`
	Errors    []string `json:"errors"`
	Timestamp int64    `json:"timestamp"`
	Code      int      `json:"code"`
	Data      T        `json:"data"`
}
