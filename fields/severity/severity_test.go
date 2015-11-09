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
	"github.com/mediaFORGE/gol/fields/severity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SeverityTestSuite struct {
	suite.Suite
}

func (s *SeverityTestSuite) TestTypeString() {
	assert.Equal(s.T(), "UNKNOWN", severity.Type(-1).String())

	assert.Equal(s.T(), "EMERGENCY", severity.Type(severity.Emergency).String())
	assert.Equal(s.T(), "ALERT", severity.Type(severity.Alert).String())
	assert.Equal(s.T(), "CRITICAL", severity.Type(severity.Critical).String())
	assert.Equal(s.T(), "ERROR", severity.Type(severity.Error).String())
	assert.Equal(s.T(), "WARNING", severity.Type(severity.Warning).String())
	assert.Equal(s.T(), "NOTICE", severity.Type(severity.Notice).String())
	assert.Equal(s.T(), "INFO", severity.Type(severity.Info).String())
	assert.Equal(s.T(), "DEBUG", severity.Type(severity.Debug).String())

	assert.Equal(s.T(), "UNKNOWN", severity.Type(8).String())
}

func (s *SeverityTestSuite) TestTypeValidate() {
	assert.NotNil(s.T(), severity.Type(-1).Validate())

	for i := 0; i < 8; i++ {
		assert.Nil(s.T(), severity.Type(i).Validate())
	}

	assert.NotNil(s.T(), severity.Type(8).Validate())
}
