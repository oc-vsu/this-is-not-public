package mocking

import (
	"testing"

	repoMock "opitz-consulting.com/examples/unittesting/mocking/internal/app/mocking/repository/mock"

	"github.com/stretchr/testify/assert"
)

// TestLogic_DoSomething uses the prefix "Test" combined with the name of the struct the
// function to be tested is bound on.
// This prefix is concat with an underscore to the function name.
func TestLogic_DoSomething(t *testing.T) {
	// we declare a new variable with our mocking struct (mocked repository) as it's value
	mockedRepo := repoMock.Memory{}

	// pass this variable into the "constructor" of our logic
	logic := NewLogic(&mockedRepo)

	err := logic.DoSomething()

	assert.NoError(t, err)
}
