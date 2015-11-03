package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// TestSuite runs the test suite.
func TestSuite(t *testing.T) {
	suite.Run(t, new(BaseLoggerTestSuite))
	suite.Run(t, new(LogMessageTestSuite))
}
