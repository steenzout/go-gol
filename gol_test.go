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

package gol_test

import (
	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/fields"
	"github.com/mediaFORGE/gol/fields/severity"
	"github.com/mediaFORGE/gol/internal/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BaseLoggerTestSuite struct {
	suite.Suite
}

type setupBaseLoggerTest struct {
	setUp func(
		msg *gol.LogMessage, mf *mock.MockLogFilter, mfmt *mock.MockLogFormatter, mw *mock.MockWriter,
	) *gol.BaseLogger
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

	in := map[string]setupBaseLoggerTest{
		"error": setupBaseLoggerTest{
			setUp: func(
				msg *gol.LogMessage, mf *mock.MockLogFilter, mfmt *mock.MockLogFormatter, mw *mock.MockWriter,
			) (logger *gol.BaseLogger) {
				mf.Mock.On("Filter", msg).Return(false, nil)
				mfmt.Mock.On("Format", msg).Return("ERROR", nil)
				mw.Mock.On("Write", []byte("ERROR")).Return(5, nil)

				logger = &gol.BaseLogger{}
				logger.SetFilter(mf)
				logger.SetFormatter(mfmt)
				logger.SetWriter(mw)

				return
			},
			message: map[string]interface{}{
				fields.Severity: severity.Error,
			},
			output: "ERROR",
		},
		"info": setupBaseLoggerTest{
			setUp: func(
				msg *gol.LogMessage, mf *mock.MockLogFilter, mfmt *mock.MockLogFormatter, mw *mock.MockWriter,
			) (logger *gol.BaseLogger) {
				mf.Mock.On("Filter", msg).Return(true, nil)

				logger = &gol.BaseLogger{}
				logger.SetFilter(mf)
				logger.SetFormatter(mfmt)
				logger.SetWriter(mw)

				return
			},
			message: map[string]interface{}{
				fields.Severity: severity.Info,
			},
			output: "",
		},
	}

	for _, t := range in {
		mf := &mock.MockLogFilter{}
		mfmt := &mock.MockLogFormatter{}
		mw := &mock.MockWriter{}
		logger := t.setUp(&t.message, mf, mfmt, mw)

		logger.Send(&t.message)

		mf.AssertExpectations(s.T())
		mfmt.AssertExpectations(s.T())
		mw.AssertExpectations(s.T())
	}
}
