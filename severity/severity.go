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

package messages

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
