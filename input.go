package main

import (
	"io"
	"io/ioutil"
)

type IInput io.Reader

type MInput struct {
	IInput
	message string
}

func (i MInput) GetMessage() (string, error) {
	message, err := ioutil.ReadAll(i)
	i.message = string(message)
	return i.message, err
}
