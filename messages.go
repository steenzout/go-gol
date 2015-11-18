package gol

import (
	"fmt"
	"time"

	"github.com/mediaFORGE/gol/fields"
	field_severity "github.com/mediaFORGE/gol/fields/severity"
	field_timestamp "github.com/mediaFORGE/gol/fields/timestamp"
)

// LogMessage is a log message.
type LogMessage map[string]interface{}

// FieldLength returns the number of fields in the message.
func (msg LogMessage) FieldLength() (n int) {
	return len(msg)
}

// Get returns the value of the given logger message field.
func (msg LogMessage) Get(f string) (i interface{}, err error) {
	if v, ok := msg[f]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("Message does not contain field %s", f)
}

// Severity returns the value of the logger message severity level field.
func (msg LogMessage) Severity() (lvl field_severity.Type, err error) {
	var v interface{}
	if v, err = msg.Get(fields.Severity); err == nil {
		return v.(field_severity.Type), nil
	}
	return field_severity.Type(-1), err
}

// SetSeverity sets the value of the logger message severity level field.
func (msg LogMessage) SetSeverity(lvl field_severity.Type) (err error) {
	if err = lvl.Validate(); err == nil {
		msg[fields.Severity] = lvl
	}
	return
}

// Start returns the value of the logger message start field.
func (msg LogMessage) Start() (s *time.Time, err error) {
	var v interface{}
	if v, err = msg.Get(fields.Start); err == nil {
		return v.(*time.Time), nil
	}
	return &time.Time{}, err
}

// SetStart sets the value of the logger message start field.
func (msg LogMessage) SetStart(s *time.Time) (err error) {
	msg[fields.Start] = s
	return nil
}

// Stop returns the value of the logger message stop field.
func (msg LogMessage) Stop() (s *time.Time, err error) {
	var v interface{}
	if v, err = msg.Get(fields.Stop); err == nil {
		return v.(*time.Time), nil
	}
	return &time.Time{}, err
}

// SetStop sets the value of the logger message stop field.
func (msg LogMessage) SetStop(s *time.Time) (err error) {
	msg[fields.Stop] = s
	return nil
}

// Timestamp returns the value of the logger message timestamp field.
func (msg LogMessage) Timestamp() (*field_timestamp.Type, error) {
	var v interface{}
	var err error
	if v, err = msg.Get(fields.Timestamp); err == nil {
		vt := v.(field_timestamp.Type)
		return &vt, nil
	}
	return nil, err
}

// NewLogMessageFunc is the function signature of LogMessage constructor functions.
type NewLogMessageFunc func(args ...interface{}) *LogMessage

// NewEmergency builds an emergency severity message.
func NewEmergency(args ...interface{}) *LogMessage {
	return NewMessage(field_severity.Type(field_severity.Emergency), args...)
}

// NewAlert builds an alert severity message.
func NewAlert(args ...interface{}) *LogMessage {
	return NewMessage(field_severity.Type(field_severity.Alert), args...)
}

// NewCritical builds a critical severity message.
func NewCritical(args ...interface{}) *LogMessage {
	return NewMessage(field_severity.Type(field_severity.Critical), args...)
}

// NewError builds an error severity message.
func NewError(args ...interface{}) *LogMessage {
	return NewMessage(field_severity.Type(field_severity.Error), args...)
}

// NewWarning builds a warning severity message.
func NewWarning(args ...interface{}) *LogMessage {
	return NewMessage(field_severity.Type(field_severity.Warning), args...)
}

// NewNotice builds a notice severity message.
func NewNotice(args ...interface{}) *LogMessage {
	return NewMessage(field_severity.Type(field_severity.Notice), args...)
}

// NewInfo builds an info severity message.
func NewInfo(args ...interface{}) *LogMessage {
	return NewMessage(field_severity.Type(field_severity.Info), args...)
}

// NewDebug builds a debug severity message.
func NewDebug(args ...interface{}) *LogMessage {
	return NewMessage(field_severity.Type(field_severity.Debug), args...)
}

// NewMessage build a log message with the given severity level.
func NewMessage(l field_severity.Type, args ...interface{}) *LogMessage {
	msg := LogMessage{
		fields.Timestamp: field_timestamp.Type{time.Now()},
		fields.Severity:  l,
	}
	for i := 0; i < len(args); i += 2 {
		msg[args[i].(string)] = args[i+1]
	}

	return &msg
}
