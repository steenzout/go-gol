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

package formatters_test

import (
	"github.com/mediaFORGE/gol"

	"github.com/mediaFORGE/gol/formatters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JSONTestSuite struct {
	suite.Suite
}

func (s *JSONTestSuite) TestFormat() {
	msg := &gol.LogMessage{
		"key1": "value1",
		"key2": "value2",
	}
	fmt := &formatters.JSON{}

	o, err := fmt.Format(msg)
	assert.True(
		s.T(),
		"{\"key1\":\"value1\",\"key2\":\"value2\"}\n" == o,
		"{\"key2\":\"value2\",\"key1\":\"value1\"}\n" == o)
	assert.Nil(s.T(), err)
}
