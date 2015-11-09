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
	"encoding/json"
	"fmt"

	"github.com/mediaFORGE/gol"
)

// JSON struct for a generic JSON formatter.
type JSON struct{}

// JSON returns a JSON representation of this struct.
func (f JSON) Format(msg *gol.LogMessage) (string, error) {
	byteArr, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\n", string(byteArr)), nil
}

var _ gol.LogFormatter = (*JSON)(nil)
