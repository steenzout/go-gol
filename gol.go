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

package gol

import (
	"io"
)

// LogFilter the interface a log filter needs to implement.
type LogFilter interface {
	Filter(*LogMessage) bool
}

// LogFormatter the interface a log message formatter needs to implement.
type LogFormatter interface {
	Format(*LogMessage) (string, error)
}

// Logger the interface a log message consumer must implement.
type Logger interface {
	Close()
	Filter() LogFilter
	Formatter() LogFormatter
	Run(chan *LogMessage)
	Send(*LogMessage) error
	SetFilter(LogFilter) error
	SetFormatter(LogFormatter) error
	SetWriter(io.Writer) error
	Status() bool
	Writer() io.Writer
}

// LoggerManager the interface to manage an application set of loggers.
type LoggerManager interface {
	Close()
	Deregister(n string) error
	Disable(n string) error
	Enable(n string) error
	IsEnabled(n string) (bool, error)
	List() []string
	Register(n string, l Logger) error
	Run(chan *LogMessage)
	Send(*LogMessage) (err error)
}

// Manager is the instance responsible for handling log messages and sending them to all registered loggers.
var Manager LoggerManager
