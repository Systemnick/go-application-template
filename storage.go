package main

import "fmt"

type IStorage interface {
	FindRecord(key string) (string, error)
	SaveRecord(key string, value string) error
}

type Storage struct {
	Dsn string
}

func (s Storage) FindRecord(key string) (string, error) {

	return "", fmt.Errorf("not implemented yet")
}

func (s Storage) SaveRecord(key string, value string) error {

	return fmt.Errorf("not implemented yet")
}
