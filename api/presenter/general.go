package presenter

import "github.com/gofiber/fiber/v2"

type DataMessage struct {
	Message    string
	SubMessage string
}

func SuccessRes(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func FailedRes(msg string, err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   msg,
		"error":  err.Error(),
	}
}

func FailedWithBodyRes(msg string, subMsg string, err error) *fiber.Map {
	data := DataMessage{
		Message:    msg,
		SubMessage: subMsg,
	}
	return &fiber.Map{
		"status": false,
		"data":   data,
		"error":  err.Error(),
	}
}
