package common

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type ApiResponse[T any] struct {
	Success   bool     `json:"success"`
	Message   string   `json:"message"`
	Errors    []string `json:"errors"`
	Timestamp int64    `json:"timestamp"`
	Code      int      `json:"code"`
	Data      T        `json:"data"`
}

func NewApiResponse() *ApiResponse[any] {
	return &ApiResponse[any]{
		Success:   true,
		Data:      nil,
		Errors:    []string{},
		Timestamp: time.Now().UnixMilli(),
		Code:      fiber.StatusOK,
	}
}
