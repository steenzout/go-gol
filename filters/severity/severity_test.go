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

package severity_test

import (
	"github.com/mediaFORGE/gol"
	field "github.com/mediaFORGE/gol/fields/severity"
	filter "github.com/mediaFORGE/gol/filters/severity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FilterTestSuite struct {
	suite.Suite
}

func (s *FilterTestSuite) TestFilter() {
	f := filter.New(field.Type(field.Emergency))
	cases := []gol.NewLogMessageFunc{
		gol.NewAlert, gol.NewCritical, gol.NewError, gol.NewWarning, gol.NewNotice, gol.NewInfo, gol.NewDebug,
	}

	assert.False(s.T(), f.Filter(gol.NewEmergency()))
	for _, newFunc := range cases {
		assert.True(s.T(), f.Filter(newFunc()))
	}

	f = filter.New(field.Type(field.Debug))
	cases = []gol.NewLogMessageFunc{
		gol.NewEmergency, gol.NewAlert, gol.NewCritical, gol.NewError, gol.NewWarning, gol.NewNotice, gol.NewInfo,
	}

	for _, newFunc := range cases {
		assert.False(s.T(), f.Filter(newFunc()))
	}
	assert.False(s.T(), f.Filter(gol.NewDebug()))
}

func (s *FilterTestSuite) TestFilterNoSeverityField() {
	f := filter.New(field.Type(field.Debug))
	assert.True(s.T(), f.Filter(&gol.LogMessage{}))
}
