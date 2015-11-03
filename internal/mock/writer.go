package mock

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockWriter struct {
	mock.Mock
}

func (m *MockWriter) Write(p []byte) (n int, err error) {
	args := m.Mock.Called(p)

	return args.Get(0).(int), args.Get(1).(error)
}

var _ io.Writer = (*MockWriter)(nil)
