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

package threshold_test

import (
	"time"

	"github.com/mediaFORGE/gol"
	fields "github.com/mediaFORGE/gol/fields"
	filter "github.com/mediaFORGE/gol/filters/threshold"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FilterTestSuite struct {
	suite.Suite
	now time.Time
}

func (s *FilterTestSuite) SetupTest() {
	s.now = time.Now()
}

func (s *FilterTestSuite) TestFilter() {
	f := filter.New(2 * time.Second)
	t1sago := s.now.Add(-1 * time.Second)
	t2sago := s.now.Add(-2 * time.Second)
	t3sago := s.now.Add(-3 * time.Second)

	assert.False(s.T(), f.Filter(&gol.LogMessage{
		fields.Start: &t1sago,
		fields.Stop:  &s.now,
	}))
	assert.True(s.T(), f.Filter(&gol.LogMessage{
		fields.Start: &t2sago,
		fields.Stop:  &s.now,
	}))
	assert.True(s.T(), f.Filter(&gol.LogMessage{
		fields.Start: &t3sago,
		fields.Stop:  &s.now,
	}))
}

func (s *FilterTestSuite) TestFilterNoStart() {
	f := filter.New(1 * time.Second)
	assert.True(s.T(), f.Filter(&gol.LogMessage{
		fields.Stop: &s.now,
	}))
}

func (s *FilterTestSuite) TestFilterNoStop() {
	f := filter.New(1 * time.Second)
	assert.True(s.T(), f.Filter(&gol.LogMessage{
		fields.Start: &s.now,
	}))
}
