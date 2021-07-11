package main

import "fmt"

type ErrorCode string

const (
	ErrorNotFound     ErrorCode = "ErrorNotFound"
	ErrorBadInput     ErrorCode = "ErrorBadInput"
	ErrorUnauthorized ErrorCode = "ErrorUnauthorized"
)

type Error struct {
	msg  string
	code ErrorCode
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code %s: %s", e.code, e.msg)
}
