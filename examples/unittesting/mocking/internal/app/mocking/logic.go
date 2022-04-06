package mocking

import (
	"opitz-consulting.com/examples/unittesting/mocking/internal/app/mocking/repository"
)

type Logic struct {
	// using the interface as type, so we could use every struct that implement this interface for our logic
	MockingRepo repository.Repository
}

// NewLogic the "constructor" expects also the interface type for our repository.
func NewLogic(m repository.Repository) *Logic {
	return &Logic{
		MockingRepo: m,
	}
}

func (l *Logic) DoSomething() error {
	var x interface{}

	err := l.MockingRepo.Insert(x)
	if err != nil {
		return err
	}

	_, err = l.MockingRepo.FindOne("something")
	if err != nil {
		return err
	}

	_, err = l.MockingRepo.Find()
	if err != nil {
		return err
	}

	return nil
}
