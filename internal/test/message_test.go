package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/mediaFORGE/gol"
)

type LogMessageTestSuite struct {
	suite.Suite
}

func (s *LogMessageTestSuite) TestStruct() {
	rec := gol.LogMessage{
		"key": "value",
	}

	assert.Equal(s.T(), rec["key"], "value")
}
