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

package simple

import (
	"fmt"
	"io"
	"sync"

	"github.com/mediaFORGE/gol"
)

// Logger generic struct for a logger.
type Logger struct {
	channel   chan *gol.LogMessage
	mutex     sync.Mutex
	waitGroup sync.WaitGroup
	status    bool
	filter    gol.LogFilter
	formatter gol.LogFormatter
	writer    io.Writer
}

// New creates and initializes a generic logger struct.
func New(f gol.LogFilter, fmt gol.LogFormatter, w io.Writer) *Logger {
	return &Logger{
		mutex:     sync.Mutex{},
		waitGroup: sync.WaitGroup{},
		status:    false,
		filter:    f,
		formatter: fmt,
		writer:    w,
	}
}

// Close closes the log message channel and waits for all processing to complete.
func (l *Logger) Close() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.status {
		l.waitGroup.Wait()
		l.status = false
	}
}

// process processes messages from logger channel.
func (l *Logger) process() {
	for msg := range l.channel {
		l.Send(msg)
	}
	l.waitGroup.Done()
}

// Filter returns the logger filter.
func (l *Logger) Filter() gol.LogFilter {
	return l.filter
}

// Formatter returns the logger formatter.
func (l *Logger) Formatter() gol.LogFormatter {
	return l.formatter
}

// Run runs a go routine to process messages from the logger channel.
func (l *Logger) Run(c chan *gol.LogMessage) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if !l.status {
		l.waitGroup.Add(1)
		l.channel = c

		go l.process()
		l.status = true
	}
}

// Send process log message.
func (l *Logger) Send(m *gol.LogMessage) (err error) {
	if m == nil {
		return
	}

	if l.Filter() != nil && l.filter.Filter(m) {
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
func (l *Logger) SetFilter(f gol.LogFilter) error {
	l.filter = f
	return nil
}

// SetFormatter sets the logger formatter.
func (l *Logger) SetFormatter(f gol.LogFormatter) error {
	if f == nil {
		return fmt.Errorf("Nil formatter")
	}
	l.formatter = f
	return nil
}

// SetWriter sets the logger writer.
func (l *Logger) SetWriter(w io.Writer) error {
	if w == nil {
		return fmt.Errorf("Nil writer")
	}
	l.writer = w
	return nil
}

// Status returns the logger running status.
// True means the logger go routine is running; False otherwise.
func (l *Logger) Status() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.status
}

// Writer returns the logger writer.
func (l *Logger) Writer() io.Writer {
	return l.writer
}

var _ gol.Logger = (*Logger)(nil)
