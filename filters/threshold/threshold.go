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

package threshold

import (
	"time"

	"github.com/mediaFORGE/gol"
)

// Threshold the struct for the filter.
type Threshold struct {
	threshold time.Duration
}

// New creates a threshold filter.
func New(d time.Duration) gol.LogFilter {
	return &Threshold{
		threshold: d,
	}
}

// Filter performs a filter check on the given message.
// Returns whether or not a given message should be filtered.
func (f Threshold) Filter(msg *gol.LogMessage) bool {
	var start *time.Time
	var stop *time.Time
	var err error

	if start, err = msg.Start(); err != nil {
		return true
	}
	if stop, err = msg.Stop(); err != nil {
		return true
	}

	return stop.Sub(*start) >= f.threshold
}

var _ gol.LogFilter = (*Threshold)(nil)
