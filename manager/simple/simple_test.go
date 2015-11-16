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

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/loggers/mock"
	"github.com/mediaFORGE/gol/manager/simple"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ManagerTestSuite struct {
	suite.Suite
	manager gol.LoggerManager
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
	s.manager = simple.New()
}

func (s *ManagerTestSuite) TestDeregister() {
	// setup
	logger := mock.New()
	s.manager.Register("mock", logger)

	// deregister
	assert.Nil(s.T(), s.manager.Deregister("mock"))
	assert.Equal(s.T(), []string{}, s.manager.List())

	// inexistent
	assert.NotNil(s.T(), s.manager.Deregister("inexistent"))
}

func (s *ManagerTestSuite) TestDisable() {
	// setup
	logger := mock.New()
	s.manager.Register("mock", logger)

	// disable
	assert.Nil(s.T(), s.manager.Disable("mock"))
	s.testIsEnabled("mock", false, nil)

	// inexistent
	assert.NotNil(s.T(), s.manager.Disable("inexistent"))
}

func (s *ManagerTestSuite) TestEnable() {
	// setup
	logger := mock.New()
	s.manager.Register("mock", logger)

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
	logger := mock.New()
	s.manager.Register("mock", logger)

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

	logger := mock.New()
	assert.Nil(s.T(), s.manager.Register("mock", logger))
	assert.Equal(s.T(), []string{"mock"}, s.manager.List())
}

func (s *ManagerTestSuite) TestRegister() {
	logger := mock.New()
	assert.Nil(s.T(), s.manager.Register("mock", logger))
	assert.Equal(s.T(), []string{"mock"}, s.manager.List())
	s.testIsEnabled("mock", true, nil)

	// duplicate
	assert.Nil(s.T(), s.manager.Register("mock", logger))
	assert.Equal(s.T(), []string{"mock"}, s.manager.List())
	s.testIsEnabled("mock", true, nil)

	// nil
	assert.NotNil(s.T(), s.manager.Register("mock", nil))
}
