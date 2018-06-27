package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

type customError struct {
	err     error
	ID      int
	Message string
}

func (p *customError) Error() string {
	return fmt.Sprintf("Error ID %d and message: %s", p.ID, p.Message)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	log.Print("1. Statement log.Print()")
	log.Printf("2. statement log.Prinf(%s)", "param")
	log.Println("3. statement log.Println()")

	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")
	fmt.Print(&buf)

	if err := raiseCustomErr(); err != nil {
		log.Printf("5. custom error %v", err)
	}

	log.Panicf("6. statement panic", raiseCustomErr)
}

func raiseCustomErr() error {
	ctx := errors.New("TestErr")
	return &customError{err: ctx, ID: 1, Message: "Custom error msg"}
}
