package api_errors

import (
	"fmt"
	"log"
)

type DbError struct {
	Func   string
	Reason string
	Err    error
	What   string
}

func (e *DbError) Format() string {
	e.What = fmt.Sprintf("At %s -> %v", e.Func, e.Err)
	return e.What
}

func (e *DbError) Log(msg string) {
	log.Fatalf("%s: %v", msg, e.Format())
}
