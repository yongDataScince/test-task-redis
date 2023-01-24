package utils

import (
	"fmt"
	"time"
)

func FormatErrorMsg(msg string, err error) string {
	return fmt.Sprintf("Message: %s\nError: %s", msg, err.Error())
}

func GetCurrDate() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d", now.Local().Month(), now.Year())
}
