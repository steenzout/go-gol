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
	//"time"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/formatters"

	field_severity "github.com/mediaFORGE/gol/fields/severity"
	filter_severity "github.com/mediaFORGE/gol/filters/severity"
	logger_simple "github.com/mediaFORGE/gol/loggers/simple"
	manager_simple "github.com/mediaFORGE/gol/manager/simple"
)

// LogMessageChannelCapacity the capacity of the log message channel.
const LogMessageChannelCapacity = 1024

// Log holds the application LogManager instance.
var Log gol.LoggerManager

func init() {
	fmt.Println("init():start")
	Log = manager_simple.New(LogMessageChannelCapacity)

	f := filter_severity.New(field_severity.Info)
	formatter := formatters.Text{}
	logger := logger_simple.New(f, formatter, os.Stdout)
	Log.Register("main", logger)

	Log.Run()
	Log.Send(gol.NewInfo("message", "main.Log has been configured"))
	fmt.Println("init():end")
}

func main() {
	fmt.Println("Started application...")
	for i := 0; i < 2000; i++ {
		Log.Send(gol.NewInfo("i", fmt.Sprintf("%d", i)))
		//time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Ending application...")
}
