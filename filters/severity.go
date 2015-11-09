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

package filters

import (
	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/fields/severity"
)

type Severity struct {
	minimum severity.Type
}

func NewSeverity(s severity.Type) gol.LogFilter {
	return &Severity{
		minimum: s,
	}
}

func (f Severity) Filter(msg *gol.LogMessage) bool {

	if s, err := msg.Severity(); err != nil {
		// no severity
		return true
	} else {
		return s > f.minimum
	}
}

var _ gol.LogFilter = (*Severity)(nil)
