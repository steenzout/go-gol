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

package simple_test

import (
	"fmt"
	"time"

	"github.com/mediaFORGE/gol"
	mfmock "github.com/mediaFORGE/gol/internal/mock"
	logger_mock "github.com/mediaFORGE/gol/loggers/mock"
	logger_simple "github.com/mediaFORGE/gol/loggers/simple"
	manager_simple "github.com/mediaFORGE/gol/managers/simple"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ManagerTestSuite struct {
	suite.Suite
	manager gol.LoggerManager
	channel chan *gol.LogMessage
}

func (s *ManagerTestSuite) testIsEnabled(n string, b bool, e error) {

	status, err := s.manager.IsEnabled(n)
	if e == nil {
		assert.Equal(s.T(), b, status)
		assert.Nil(s.T(), err)
	} else {
		assert.False(s.T(), status)
		assert.NotNil(s.T(), err)
	}
}

func (s *ManagerTestSuite) SetupTest() {
	s.manager = manager_simple.New()
	s.channel = make(chan *gol.LogMessage, 1)
}

func (s *ManagerTestSuite) TeardownTest() {
	s.manager.Close()
	close(s.channel)
}

func (s *ManagerTestSuite) TestDeregister() {
	// setup
	l := logger_mock.New()
	s.manager.Register("mock", l)

	// deregister
	assert.Nil(s.T(), s.manager.Deregister("mock"))
	assert.Equal(s.T(), []string{}, s.manager.List())

	// inexistent
	assert.NotNil(s.T(), s.manager.Deregister("inexistent"))
}

func (s *ManagerTestSuite) TestDisable() {
	// setup
	l := logger_mock.New()
	s.manager.Register("mock", l)

	// disable
	assert.Nil(s.T(), s.manager.Disable("mock"))
	s.testIsEnabled("mock", false, nil)

	// inexistent
	assert.NotNil(s.T(), s.manager.Disable("inexistent"))
}

func (s *ManagerTestSuite) TestEnable() {
	// setup
	l := logger_mock.New()
	s.manager.Register("mock", l)

	// registered logger is enabled by default
	s.testIsEnabled("mock", true, nil)

	// enable a disabled logger
	s.manager.Disable("mock")
	assert.Nil(s.T(), s.manager.Enable("mock"))
	s.testIsEnabled("mock", true, nil)

	// inexistent
	assert.NotNil(s.T(), s.manager.Enable("inexistent"))
}

func (s *ManagerTestSuite) TestIsEnabled() {
	// setup
	l := logger_mock.New()
	s.manager.Register("mock", l)

	// enabled logger
	s.testIsEnabled("mock", true, nil)

	// disabled logger
	s.manager.Disable("mock")
	s.testIsEnabled("mock", false, nil)

	// inexistent logger
	s.testIsEnabled("inexistent", false, fmt.Errorf("error"))
}

func (s *ManagerTestSuite) TestList() {
	assert.Equal(s.T(), []string{}, s.manager.List())

	l := logger_mock.New()
	assert.Nil(s.T(), s.manager.Register("mock", l))
	assert.Equal(s.T(), []string{"mock"}, s.manager.List())
}

func (s *ManagerTestSuite) TestRegister() {
	l := logger_mock.New()
	assert.Nil(s.T(), s.manager.Register("mock", l))
	assert.Equal(s.T(), []string{"mock"}, s.manager.List())
	s.testIsEnabled("mock", true, nil)

	// duplicate
	assert.Nil(s.T(), s.manager.Register("mock", l))
	assert.Equal(s.T(), []string{"mock"}, s.manager.List())
	s.testIsEnabled("mock", true, nil)

	// nil
	assert.NotNil(s.T(), s.manager.Register("mock", nil))
}

func (s *ManagerTestSuite) TestSend() {
	m := gol.NewEmergency("field", "value")

	// l1 will not filter the message
	mf1 := &mfmock.LogFilter{}
	mf1.On("Filter", m).Return(false)
	mfmt1 := &mfmock.LogFormatter{}
	mfmt1.On("Format", m).Return("EMERGENCY field=value", nil)
	mw1 := &mfmock.Writer{}
	mw1.On("Write", mock.Anything).Return(21, nil)
	l1 := logger_simple.New(mf1, mfmt1, mw1)

	// l2 will filter the message
	mf2 := &mfmock.LogFilter{}
	mf2.On("Filter", m).Return(true)
	mfmt2 := &mfmock.LogFormatter{}
	mw2 := &mfmock.Writer{}
	l2 := logger_simple.New(mf2, mfmt2, mw2)

	s.manager.Register("l1", l1)
	s.manager.Register("l2", l2)

	s.manager.Run(s.channel)

	assert.Nil(s.T(), s.manager.Send(m))
	time.Sleep(1 * time.Second)
	s.manager.Close()

	mf1.AssertExpectations(s.T())
	mfmt1.AssertExpectations(s.T())
	mw1.AssertExpectations(s.T())

	mf2.AssertExpectations(s.T())
	mfmt2.AssertExpectations(s.T())
	mw2.AssertExpectations(s.T())
}

func (s *ManagerTestSuite) TestSendWithoutRun() {
	m := gol.NewEmergency("field", "value")

	assert.Equal(s.T(), s.manager.Send(m), fmt.Errorf("manager.simple.LogManager is not running"))
}
