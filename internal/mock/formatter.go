package mock

import (
	"github.com/mediaFORGE/gol"

	"github.com/stretchr/testify/mock"
)

type MockLogFormatter struct {
	mock.Mock
}

func (m *MockLogFormatter) Format(msg gol.LogMessage) (string, error) {
	args := m.Mock.Called(msg)

	return args.Get(0).(string), args.Get(1).(error)
}

var _ gol.LogFormatter = (*MockLogFormatter)(nil)

