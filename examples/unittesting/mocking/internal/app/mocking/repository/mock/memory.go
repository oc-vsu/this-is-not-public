package mock

import (
	"errors"
)

// Memory is our testing struct
// It has all methods to implement the repository.Repository interface.
// Instead of connection to a database service, it reads / writes to the memory.
type Memory struct{}

var memory []interface{}

func (m *Memory) Insert(x interface{}) error {
	memory = append(memory, x)

	return nil
}

func (m *Memory) FindOne(id string) (interface{}, error) {
	for _, x := range memory {
		if x == id {
			return x, nil
		}
	}

	return nil, errors.New("not found")
}

func (m *Memory) Find() ([]interface{}, error) {
	return memory, nil
}
