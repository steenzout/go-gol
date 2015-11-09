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

// Log base struct of a logger.
type Log struct {
	filter    LogFilter
	formatter LogFormatter
	writer    io.Writer
}

// SimpleLog creates and initializes a logger struct.
func SimpleLog(f LogFilter, fmt LogFormatter, w io.Writer) *Log {
	return &Log{f, fmt, w}
}

// Filter returns the logger filter.
func (l *Log) Filter() LogFilter {
	return l.filter
}

// Formatter returns the logger formatter.
func (l *Log) Formatter() LogFormatter {
	return l.formatter
}

// Writer returns the logger writer.
func (l *Log) Writer() io.Writer {
	return l.writer
}

// Send process log message.
func (l *Log) Send(m *LogMessage) (err error) {
	if m == nil {
		return
	}

	if l.filter != nil && l.filter.Filter(m) {
		return
	}

	if l.formatter == nil {
		return fmt.Errorf("log formatter is nil")
	}

	var msg string
	if msg, err = l.formatter.Format(m); err != nil {
		return
	}

	if l.writer == nil {
		return fmt.Errorf("log writer is nil")
	}

	_, err = l.writer.Write([]byte(msg))
	return
}

// SetFilter sets the logger filter.
func (l *Log) SetFilter(f LogFilter) error {
	l.filter = f
	return nil
}

// SetFormatter sets the logger formatter.
func (l *Log) SetFormatter(f LogFormatter) error {
	if f == nil {
		return fmt.Errorf("Nil formatter")
	}
	l.formatter = f
	return nil
}

// SetWriter sets the logger writer.
func (l *Log) SetWriter(w io.Writer) error {
	if w == nil {
		return fmt.Errorf("Nil writer")
	}
	l.writer = w
	return nil
}

var _ Logger = (*Log)(nil)
