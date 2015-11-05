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

package gol

import (
	"fmt"

	"github.com/mediaFORGE/gol/fields/severity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MessageTestSuite struct {
	suite.Suite
}

func (s *MessageTestSuite) TestGet() {
	msg := LogMessage{
		"key": "value",
	}

	assert.Equal(s.T(), msg["key"], "value")

	v, err := msg.Get("key")
	assert.Equal(s.T(), "value", v)
	assert.Nil(s.T(), err)

	v, err = msg.Get("unknown")
	assert.Nil(s.T(), v)
	assert.Equal(s.T(), fmt.Errorf("Message does not contain field unknown"), err)
}

func (s *MessageTestSuite) TestGetSetSeverity() {
	msg := LogMessage{
		"key": "value",
	}

	v, err := msg.GetSeverity()
	assert.Equal(s.T(), severity.Type(-1), v)
	assert.Equal(s.T(), fmt.Errorf("Message does not contain field severity"), err)

	lvl := severity.Type(severity.Emergency)
	msg.SetSeverity(lvl)

	v, err = msg.GetSeverity()
	assert.Equal(s.T(), lvl, v)
	assert.Nil(s.T(), err)
}

func (s *MessageTestSuite) assertSeverityLevel(expected severity.Type, f NewLogMessageFunc) {
	msg := f()
	severity, err := msg.GetSeverity()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expected, severity)
}

func (s *MessageTestSuite) TestNewSeverity() {
	cases := map[int]NewLogMessageFunc{
		severity.Emergency: NewEmergency,
		severity.Alert:     NewAlert,
		severity.Critical:  NewCritical,
		severity.Error:     NewError,
		severity.Warning:   NewWarning,
		severity.Notice:    NewNotice,
		severity.Info:      NewInfo,
		severity.Debug:     NewDebug,
	}

	for lvl, f := range cases {
		s.assertSeverityLevel(severity.Type(lvl), f)
	}
}
