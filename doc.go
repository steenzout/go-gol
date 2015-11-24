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

/*
Package gol is a library that enables an application for perform logging operations asynchronously.

To achieve this goal and give flexibility it provides several logging interfaces:
- Logger to implement a logger
- LoggerFilter to implement a log message filter
- LogFormatter to implement a log message format
- LoggerManager to implement log management

It also provides several types, namely:
- LogMessage
- field.severity.Type
- field.timestamp.Type

With this in place you can create your own loggers, filters, format logic and logger management.


LogMessage

The message is a hash where the keys are the names of the message fields and
its values are the contents for those fields.

Some fields have pre-defined names which you can find at
fields.fields.go.

Types have also been added to be able to populate them with helper functions to,
for example, easily convert them to string or perform validation.
You can find these at:
- fields.severity.severity.go
- fields.timestamp.timestamp.go


Logger

Logger is typically composed of a filter, formatter and a writer.

A simple implementation is provided with this library.
In this implementation the logger will run the filter against the message.
If the filter returns true it will discard the message;
otherwise it will format it and send it to the configured writer.

Since this is an interface,
you can create simpler or fancier implementations depending
on your application requirements.


LoggerFilter

The filter is responsible for analyzing the contents of the message and
determining if it will be sent or not to the writer.

For example, a severity filter is provided with the library.
This filter retrieves the severity level field (more details on RFC5424) on the message and
if it's of lower severity level it will discard it.

Since this is an interface,
you can create simpler or fancier implementations depending
on your application requirements.


LoggerFormatter

The formatter is responsible for analyzing the contents of the message and
converting it to a string.

For example, a text formatter is provided with the library.
This formatter will go through all of the fields in the message and
start building the string with field=value pairs separated by each other by white space and
ending with a newline character.

Since this is an interface,
you can create simpler or fancier implementations depending
on your application requirements.


LoggerManager

The manager is responsible for capturing messages through a channel and
delivering them to all registered and enabled loggers.

It also provides extra functionality to manage logger registration, logger status and
perform a graceful shutdown.

Since this is an interface,
you can create simpler or fancier implementations depending
on your application requirements.


Example

A simple example where you want all messages with severity INFO or higher to be written to
standard output and formatted with the generic text formatter:

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

	// Log holds the application LogManager instance.
	var Log gol.LoggerManager

	func init() {
		fmt.Println("init():start")
		Log = manager_simple.New()

		f := filter_severity.New(field_severity.Info)
		formatter := formatters.Text{}
		logger := logger_simple.New(f, formatter, os.Stdout)
		Log.Register("main", logger)

		channel := make(chan *gol.LogMessage, 10)
		Log.Run(channel)
		Log.Send(gol.NewInfo("message", "main.Log has been configured"))
		channel <- gol.NewInfo("message", "this message was sent directly to the log manager channel")

		fmt.Println("init():end")
	}

	func main() {
		fmt.Println("Started application.")
		defer func() {
			Log.Close()
			fmt.Println("Ended application.")
		}()

		// send 10,000 messages
		for i := 0; i < 10000; i++ {
			Log.Send(gol.NewInfo("i", fmt.Sprintf("%d", i)))
		}
		fmt.Println("Ending application...")
	}
*/
package gol
