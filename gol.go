package gol

import (
	"fmt"
	"io"
)

// http://tools.ietf.org/html/rfc5424
const (
	// Emergency system is unusable
	Emergency = iota
	// Alert action must be taken immediately
	Alert
	// Critical critical conditions
	Critical
	// Error error conditions
	Error
	// Warning warning conditions
	Warning
	// Notice normal but significant condition
	Notice
	// Info informational messages
	Info
	// Debug debug-level messages
	Debug
)

// LogMessage is a log message.
type LogMessage map[string]interface{}

// LogFilter the interface a log filter needs to implement.
type LogFilter interface {
	Filter(LogMessage) (bool, error)
}

// LogFormatter the interface a log message formatter needs to implement.
type LogFormatter interface {
	Format(LogMessage) (string, error)
}

// Logger the interface a log message consumer must implement.
type Logger interface {
	Filter() LogFilter
	Formatter() LogFormatter
	Writer() io.Writer
	Send(LogMessage) error
	SetFilter(LogFilter) error
	SetFormatter(LogFormatter) error
	SetWriter(io.Writer) error
}

// BaseLogger base implementation of a logger.
type BaseLogger struct {
	filter   LogFilter
	formatter   LogFormatter
	writer io.Writer
}

func (l *BaseLogger) Filter() LogFilter {
	return l.filter
}

func (l *BaseLogger) Formatter() LogFormatter {
	return l.formatter
}

func (l *BaseLogger) Writer() io.Writer {
	return l.writer
}

func (l *BaseLogger) Send(m LogMessage) (err error)  {
	var msg string

	if m == nil {
		return fmt.Errorf("")
	}

	if filter, err := l.filter.Filter(m); err != nil {
		return err
	} else if filter {
		return nil
	}

	if msg, err = l.formatter.Format(m); err != nil {
		return nil
	}

	_, err = l.writer.Write([]byte(msg))
	return
}

func (l *BaseLogger) SetFilter(f LogFilter) (err error) {
	if f == nil {
		return fmt.Errorf("")
	}

	l.filter = f
	return
}

func (l *BaseLogger) SetFormatter(f LogFormatter) (err error)  {
	if f == nil {
		return fmt.Errorf("")
	}

	l.formatter = f
	return
}

func (l *BaseLogger) SetWriter(w io.Writer) (err error)  {
	if w == nil {
		return fmt.Errorf("")
	}

	l.writer = w
	return
}

var _ Logger = (*BaseLogger)(nil)
