package input

import (
	"io"
	"io/ioutil"
)

type IInput io.Reader

type Input struct {
	IInput
	message string
}

func (i Input) GetMessage() (string, error) {
	message, err := ioutil.ReadAll(i)
	i.message = string(message)
	return i.message, err
}
