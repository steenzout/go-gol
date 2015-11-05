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
	"fmt"
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
	Filter() LogFilter
	Formatter() LogFormatter
	Writer() io.Writer
	Send(*LogMessage) error
	SetFilter(LogFilter) error
	SetFormatter(LogFormatter) error
	SetWriter(io.Writer) error
}

// BaseLogger base implementation of a logger.
type BaseLogger struct {
	filter    LogFilter
	formatter LogFormatter
	writer    io.Writer
}

// NewBaseLogger creates and initializes a BaseLogger struct.
func NewBaseLogger(f LogFilter, fmt LogFormatter, w io.Writer) (l Logger) {
	return &BaseLogger{
		filter:    f,
		formatter: fmt,
		writer:    w,
	}
}

// Filter returns the logger filter.
func (l *BaseLogger) Filter() LogFilter {
	return l.filter
}

// Formatter returns the logger formatter.
func (l *BaseLogger) Formatter() LogFormatter {
	return l.formatter
}

// Writer returns the logger writer.
func (l *BaseLogger) Writer() io.Writer {
	return l.writer
}

// Send process log message.
func (l *BaseLogger) Send(m *LogMessage) (err error) {
	if m == nil {
		return fmt.Errorf("")
	}

	if filter := l.filter.Filter(m); filter {
		return
	}

	var msg string
	if msg, err = l.formatter.Format(m); err != nil {
		return
	}

	_, err = l.writer.Write([]byte(msg))
	return
}

// SetFilter sets the logger filter.
func (l *BaseLogger) SetFilter(f LogFilter) (err error) {
	if f == nil {
		return fmt.Errorf("")
	}

	l.filter = f
	return
}

// SetFormatter sets the logger formatter.
func (l *BaseLogger) SetFormatter(f LogFormatter) (err error) {
	if f == nil {
		return fmt.Errorf("")
	}

	l.formatter = f
	return
}

// SetWriter sets the logger writer.
func (l *BaseLogger) SetWriter(w io.Writer) (err error) {
	if w == nil {
		return fmt.Errorf("")
	}

	l.writer = w
	return
}

var _ Logger = (*BaseLogger)(nil)
