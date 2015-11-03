package mock

import (
	"github.com/mediaFORGE/gol"

	"github.com/stretchr/testify/mock"
)

type MockLogFilter struct {
	mock.Mock
}

func (m *MockLogFilter) Filter(msg gol.LogMessage) (bool, error) {
	args := m.Mock.Called(msg)

	return args.Get(0).(bool), args.Get(1).(error)
}

var _ gol.LogFilter = (*MockLogFilter)(nil)
