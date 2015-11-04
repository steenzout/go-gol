package severity

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
)

type HelpersTestSuite struct {
	suite.Suite
}

func (s *HelpersTestSuite) TestSeverityLevelString() {
	assert.Equal(s.T(), "UNKNOWN", severityLevelString(-1))

	assert.Equal(s.T(), "EMERGENCY", severityLevelString(Emergency))
	assert.Equal(s.T(), "ALERT", severityLevelString(Alert))
	assert.Equal(s.T(), "CRITICAL", severityLevelString(Critical))
	assert.Equal(s.T(), "ERROR", severityLevelString(Error))
	assert.Equal(s.T(), "WARNING", severityLevelString(Warning))
	assert.Equal(s.T(), "NOTICE", severityLevelString(Notice))
	assert.Equal(s.T(), "INFO", severityLevelString(Info))
	assert.Equal(s.T(), "DEBUG", severityLevelString(Debug))

	assert.Equal(s.T(), "UNKNOWN", severityLevelString(-1))
}

// TestSuite runs the test suites.
func TestSuite(t *testing.T) {
	suite.Run(t, new(HelpersTestSuite))
}