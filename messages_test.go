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

package gol_test

import (
	"fmt"
	"time"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/fields/severity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MessageTestSuite struct {
	suite.Suite
}

func (s *MessageTestSuite) TestFieldLength() {
	msg := gol.LogMessage{
		"key": "value",
	}
	assert.Equal(s.T(), 1, msg.FieldLength())
}

func (s *MessageTestSuite) TestGet() {
	msg := gol.LogMessage{
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
	msg := gol.LogMessage{
		"key": "value",
	}

	v, err := msg.Severity()
	assert.Equal(s.T(), severity.Type(-1), v)
	assert.Equal(s.T(), fmt.Errorf("Message does not contain field severity"), err)

	lvl := severity.Type(severity.Emergency)
	msg.SetSeverity(lvl)

	v, err = msg.Severity()
	assert.Equal(s.T(), lvl, v)
	assert.Nil(s.T(), err)
}

func (s *MessageTestSuite) TestGetSetStart() {
	msg := gol.LogMessage{
		"key": "value",
	}

	v, err := msg.Start()
	assert.Equal(s.T(), &time.Time{}, v)
	assert.Equal(s.T(), fmt.Errorf("Message does not contain field start"), err)

	start := time.Now()
	msg.SetStart(&start)

	v, err = msg.Start()
	assert.Equal(s.T(), &start, v)
	assert.Nil(s.T(), err)
}

func (s *MessageTestSuite) TestGetSetStop() {
	msg := gol.LogMessage{
		"key": "value",
	}

	v, err := msg.Stop()
	assert.Equal(s.T(), &time.Time{}, v)
	assert.Equal(s.T(), fmt.Errorf("Message does not contain field stop"), err)

	stop := time.Now()
	msg.SetStop(&stop)

	v, err = msg.Stop()
	assert.Equal(s.T(), &stop, v)
	assert.Nil(s.T(), err)
}

func (s *MessageTestSuite) assertSeverityLevel(expected severity.Type, f gol.NewLogMessageFunc) {
	msg := f("key", 1)
	severity, err := msg.Severity()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expected, severity)

	if key, err := msg.Get("key"); err != nil {
		assert.Fail(s.T(), err.Error())
	} else {
		assert.NotNil(s.T(), key)
		assert.Equal(s.T(), 1, key)
	}
}

func (s *MessageTestSuite) TestNewSeverity() {
	cases := map[severity.Type]gol.NewLogMessageFunc{
		severity.Emergency: gol.NewEmergency,
		severity.Alert:     gol.NewAlert,
		severity.Critical:  gol.NewCritical,
		severity.Error:     gol.NewError,
		severity.Warning:   gol.NewWarning,
		severity.Notice:    gol.NewNotice,
		severity.Info:      gol.NewInfo,
		severity.Debug:     gol.NewDebug,
	}

	for lvl, f := range cases {
		s.assertSeverityLevel(severity.Type(lvl), f)
	}
}

func (s *MessageTestSuite) TestTimestamp() {
	var msg *gol.LogMessage

	msg = &gol.LogMessage{"key": "value"}
	v, err := msg.Timestamp()
	assert.Equal(s.T(), fmt.Errorf("Message does not contain field timestamp"), err)
	assert.Nil(s.T(), v)

	msg = gol.NewInfo("key", "value")
	v, err = msg.Timestamp()
	assert.Nil(s.T(), err)
	assert.True(s.T(), v.Before(time.Now()))
}
