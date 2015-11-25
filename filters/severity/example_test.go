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

package severity_test

import (
	"os"

	"github.com/mediaFORGE/gol"
	field "github.com/mediaFORGE/gol/fields/severity"
	filter "github.com/mediaFORGE/gol/filters/severity"
	"github.com/mediaFORGE/gol/formatters"
	"github.com/mediaFORGE/gol/loggers/simple"
)

func ExampleSeverity() {
	txtFmt := &formatters.Text{}
	errorLog := simple.New(filter.New(field.Error), txtFmt, os.Stderr)

	errorLog.Send(gol.NewEmergency("message", "system is down"))
	// Output: EMERGENCY message:'system is down'

	errorLog.Send(gol.NewAlert("message", "failed to write to disk"))
	// Output: ALERT message:'failed to write to disk'

	errorLog.Send(gol.NewCritical("message", "high server load"))
	// Output: CRITICAL message:'high server load'

	errorLog.Send(gol.NewError("message", "invalid number format"))
	// Output: ERROR message:'invalid number format'

	errorLog.Send(gol.NewWarning("message", "performance close to 1s threshold"))
	// Output:

	errorLog.Send(gol.NewNotice("message", "failed to communicate with monitoring service"))
	// Output:

	errorLog.Send(gol.NewInfo("message", "requested processed in 250ms"))
	// Output:

	errorLog.Send(gol.NewDebug("debug", "var x = 10"))
	// Output:
}
