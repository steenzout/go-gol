package gol

import (
	"fmt"

	"github.com/mediaFORGE/gol/fields"
	"github.com/mediaFORGE/gol/fields/severity"
)

// LogMessage is a log message.
type LogMessage map[string]interface{}

// GetSeverity returns the logger Severity level.
func (msg LogMessage) Get(f string) (i interface{}, err error) {
	if v, ok := msg[f]; ok {
		fmt.Printf("%s=%s\n", f, v)
		return v, nil
	}
	fmt.Println("error")
	return nil, fmt.Errorf("Message does not contain field %s", f)
}

// GetSeverity returns the logger Severity level.
func (msg LogMessage) GetSeverity() (lvl severity.Type, err error) {
	var v interface{}
	if v, err = msg.Get(fields.Severity); err == nil {
		return v.(severity.Type), nil
	}
	return severity.Type(-1), err
}

// SetSeverity sets the logger Severity level.
func (msg LogMessage) SetSeverity(lvl severity.Type) (err error) {
	if err = lvl.Validate(); err == nil {
		msg[fields.Severity] = lvl
	}
	return
}

// NewEmergency builds an emergency severity message.
func NewEmergency(args ...interface{}) *LogMessage {
	return NewMessage(severity.Type(severity.Emergency))
}
// NewAlert builds an alert severity message.
func NewAlert(args ...interface{}) *LogMessage {
	return NewMessage(severity.Type(severity.Alert))
}

// NewCritical builds a critical severity message.
func NewCritical(args ...interface{}) *LogMessage {
	return NewMessage(severity.Type(severity.Critical))
}
// NewError builds an error severity message.
func NewError(args ...interface{}) *LogMessage {
	return NewMessage(severity.Type(severity.Error))
}

// NewWarning builds a warning severity message.
func NewWarning(args ...interface{}) *LogMessage {
	return NewMessage(severity.Type(severity.Warning))
}
// NewNotice builds a notice severity message.
func NewNotice(args ...interface{}) *LogMessage {
	return NewMessage(severity.Type(severity.Notice))
}

// NewInfo builds an info severity message.
func NewInfo(args ...interface{}) *LogMessage {
	return NewMessage(severity.Type(severity.Info))
}
// NewDebug builds a debug severity message.
func NewDebug(args ...interface{}) *LogMessage {
	return NewMessage(severity.Type(severity.Debug))
}

// NewMessage build a log message with the given severity level.
func NewMessage(l severity.Type, args ...interface{}) *LogMessage {
	return &LogMessage{
		fields.Severity: l,
	}
}
