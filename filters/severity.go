package filters

import (
	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/fields/severity"
)

type Severity struct {
	minimum severity.Type
}

func NewSeverity(s severity.Type) gol.LogFilter {
	return &Severity{
		minimum: s,
	}
}

func (f Severity) Filter(msg *gol.LogMessage) bool {
	if s, err := msg.GetSeverity(); err != nil {
		return false
	} else {
		return s <= f.minimum
	}
}
