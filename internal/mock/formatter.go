//
// Copyright 2015 Rakuten Marketing LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

