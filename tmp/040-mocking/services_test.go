package mock_demo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedCalculator struct {
	mock.Mock
}

func (m *MockedCalculator) Sum(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0)
}

func TestSum(t *testing.T) {
	// populate mock
	calc := new(MockedCalculator)
	calc.On("Sum", 10, 20).Return(30)

	//
	actual := Sum(10, 20, calc)
	assert.Equal(t, 30, actual)
	calc.AssertExpectations(t)
}
