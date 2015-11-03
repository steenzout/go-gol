package test

import (
	"io"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/internal/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BaseLoggerTestSuite struct {
	suite.Suite
}

type setupBaseLoggerTest struct {
	setUp   func(mf gol.LogFilter, mfmt gol.LogFormatter, mw io.Writer)
	message gol.LogMessage
	output  string
}

func (s *BaseLoggerTestSuite) TestGetSetFilter() {

	l := gol.BaseLogger{}

	assert.Nil(s.T(), l.Filter())
	l.SetFilter(&mock.MockLogFilter{})
	assert.NotNil(s.T(), l.Filter())
}

func (s *BaseLoggerTestSuite) TestGetSetFormatter() {

	l := gol.BaseLogger{}

	assert.Nil(s.T(), l.Formatter())
	l.SetFormatter(&mock.MockLogFormatter{})
	assert.NotNil(s.T(), l.Formatter())
}

func (s *BaseLoggerTestSuite) TestGetSetWriter() {

	l := gol.BaseLogger{}

	assert.Nil(s.T(), l.Writer())
	l.SetWriter(&mock.MockWriter{})
	assert.NotNil(s.T(), l.Writer())
}

func (s *BaseLoggerTestSuite) TestSend() {
	var in map[string]setupBaseLoggerTest

	in = map[string]setupBaseLoggerTest{
		"error": setupBaseLoggerTest{
			setUp: func(msg gol.LogMessage, mf mock.MockLogFilter, mfmt mock.MockLogFormatter, mw mock.MockWriter) {
				mf.Mock.On("Filter", msg).Return(true, nil)
				mfmt.Mock("Format", msg).Return()
			},
			message: map[string]string{
				"severity": gol.Error,
			},
			output: "ERROR",
		},
		"info": setupBaseLoggerTest{
			setUp: func(msg gol.LogMessage, mf mock.MockLogFilter, mfmt mock.MockLogFormatter, mw mock.MockWriter) {
				mf.Mock.On("Filter", msg).Return(false, nil)
			},
			message: map[string]string{
				"severity": gol.Info,
			},
			output: nil,
		},
	}



	// setUp
	mf := &mock.MockLogFilter{}


	mfmt := &mock.MockLogFormatter{}
	mw := &mock.MockWriter{}

	l := &gol.BaseLogger{}
	l.SetFilter(mf)
	l.SetFormatter(mfmt)
	l.SetWriter(mw)



}
