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

package main

import (
	"fmt"
	"os"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/formatters"

	field_severity "github.com/mediaFORGE/gol/fields/severity"
	filter_severity "github.com/mediaFORGE/gol/filters/severity"
	logger_simple "github.com/mediaFORGE/gol/loggers/simple"
	manager_simple "github.com/mediaFORGE/gol/managers/simple"
)

// LogWorkers the number of log message workers.
const LogWorkers = 4

// Log holds the application LogManager instance.
var Log gol.LoggerManager

func init() {
	fmt.Println("init():start")
	Log = manager_simple.New(LogWorkers)

	f := filter_severity.New(field_severity.Info)
	formatter := formatters.Text{}
	logger := logger_simple.New(f, formatter, os.Stdout)
	Log.Register("main", logger)

	Log.Run()
	Log.Send(gol.NewInfo("message", "main.Log has been configured"))
	fmt.Println("init():end")
}
