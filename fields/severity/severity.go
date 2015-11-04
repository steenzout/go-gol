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

package severity

import (
	"fmt"
)

// http://tools.ietf.org/html/rfc5424
const (
	// Emergency severity level for system is unusable.
	Emergency = iota
	// Alert severity level for action must be taken immediately.
	Alert
	// Critical severity level for critical conditions.
	Critical
	// Error severity level for error conditions.
	Error
	// Warning severity level for warning conditions.
	Warning
	// Notice severity level for normal but significant condition.
	Notice
	// Info severity level for informational messages.
	Info
	// Debug severity level for debug-level messages.
	Debug
)

// Type numerical value for the log message severity as specified by RFC5424.
type Type int

// String returns a string representation of the severity level.
func (lvl Type) String() string {
	switch lvl {
	case Emergency:
		return "EMERGENCY"
	case Alert:
		return "ALERT"
	case Critical:
		return "CRITICAL"
	case Error:
		return "ERROR"
	case Warning:
		return "WARNING"
	case Notice:
		return "NOTICE"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}
// String returns a string representation of the severity level.
func (lvl Type) Validate() error {
	switch lvl {
	case Emergency, Alert, Critical, Error, Warning, Notice, Info, Debug:
		return nil
	default:
		return fmt.Errorf("Invalid severity level %d", lvl)
	}
}
