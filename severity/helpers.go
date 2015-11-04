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

func severityLevelString(l int) string {
	switch l {
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
