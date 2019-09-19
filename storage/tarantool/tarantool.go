package tarantool

import (
	"fmt"

	"git.rnd.mtt/innovation/call-initiator/storage"
)

type Tarantool struct {
	storage.Storage
}

// Example:
// key = "calls.sid"
func (s Tarantool) FindRecord(key string) (string, error) {

	return "", fmt.Errorf("not implemented yet")
}

// Example:
// key = "calls"
// value = `{"sid": "", "from": "1234567"}`
func (s Tarantool) SaveRecord(key string, value string) error {

	return fmt.Errorf("not implemented yet")
}

func (s Tarantool) DeleteRecord(key string) (string, error) {

	return "", fmt.Errorf("not implemented yet")
}
