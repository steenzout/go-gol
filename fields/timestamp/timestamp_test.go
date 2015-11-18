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

package timestamp_test

import (
	"time"

	"github.com/mediaFORGE/gol/fields/timestamp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FieldTestSuite struct {
	suite.Suite
}

func (s *FieldTestSuite) TestTypeString() {
	assert.Equal(
		s.T(),
		"2015-01-02T03:04:05Z",
		timestamp.Type{
			time.Date(2015, 1, 2, 3, 4, 5, 0, time.UTC),
		}.String())
}
