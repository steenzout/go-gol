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

package formatters

import (
	"fmt"
	"strings"
	"time"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/fields"
	"github.com/mediaFORGE/gol/fields/severity"
	"github.com/mediaFORGE/gol/fields/timestamp"
)

// Text struct for a generic text formatter.
type Text struct{}

// Format formats the log message.
func (f Text) Format(msg *gol.LogMessage) (string, error) {
	lmsg := msg.FieldLength()
	buffer := make([]string, lmsg, lmsg)

	i := 0
	for k, v := range *msg {
		if k != fields.Severity && k != fields.Timestamp {
			buffer[i] = fmt.Sprintf("%s='%s'", k, v)
			i += 1
		}
	}

	var err error
	var t *timestamp.Type
	var s severity.Type
	if t, err = msg.Timestamp(); err != nil {
		t = &timestamp.Type{time.Now()}
	}

	if s, err = msg.Severity(); err != nil {
		return fmt.Sprintf("%s UNKNOWN %s\n", t.String(), strings.Join(buffer, " ")), nil
	} else {
		return fmt.Sprintf("%s %s %s\n", t.String(), s.String(), strings.Join(buffer, " ")), nil
	}
}

var _ gol.LogFormatter = (*Text)(nil)
