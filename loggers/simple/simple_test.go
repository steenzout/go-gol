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

package simple_test

import (
	"fmt"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/internal/mock"

	"github.com/mediaFORGE/gol/loggers/simple"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoggerTestSuite struct {
	suite.Suite
}

type setupLogTest struct {
	setUp func(
		msg *gol.LogMessage, mf *mock.LogFilter, mfmt *mock.LogFormatter, mw *mock.Writer,
	) *simple.Logger
	message *gol.LogMessage
	output  string
}

func (s *LoggerTestSuite) TestGetSetFilter() {

	l := simple.New(nil, nil, nil)

	assert.Nil(s.T(), l.Filter())
	assert.Nil(s.T(), l.SetFilter(&mock.LogFilter{}))
	assert.NotNil(s.T(), l.Filter())

	assert.Nil(s.T(), l.SetFilter(nil))
}

func (s *LoggerTestSuite) TestGetSetFormatter() {

	l := simple.New(nil, nil, nil)

	assert.Nil(s.T(), l.Formatter())
	assert.Nil(s.T(), l.SetFormatter(&mock.LogFormatter{}))
	assert.NotNil(s.T(), l.Formatter())

	assert.Error(s.T(), l.SetFormatter(nil))
}

func (s *LoggerTestSuite) TestGetSetWriter() {

	l := simple.New(nil, nil, nil)

	assert.Nil(s.T(), l.Writer())
	assert.Nil(s.T(), l.SetWriter(&mock.Writer{}))
	assert.NotNil(s.T(), l.Writer())

	assert.Error(s.T(), l.SetWriter(nil))
}

func (s *LoggerTestSuite) TestSend() {

	in := map[string]setupLogTest{
		"error": setupLogTest{
			setUp: func(
				msg *gol.LogMessage, mf *mock.LogFilter, mfmt *mock.LogFormatter, mw *mock.Writer,
			) (logger *simple.Logger) {
				mf.Mock.On("Filter", msg).Return(false, nil)
				mfmt.Mock.On("Format", msg).Return("ERROR", nil)
				mw.Mock.On("Write", []byte("ERROR")).Return(5, nil)

				logger = simple.New(mf, mfmt, mw)

				return
			},
			message: gol.NewError(),
			output:  "ERROR",
		},
		"info": setupLogTest{
			setUp: func(
				msg *gol.LogMessage, mf *mock.LogFilter, mfmt *mock.LogFormatter, mw *mock.Writer,
			) (logger *simple.Logger) {
				mf.Mock.On("Filter", msg).Return(true, nil)

				logger = simple.New(mf, mfmt, mw)

				return
			},
			message: gol.NewInfo(),
			output:  "",
		},
	}

	for _, t := range in {
		mf := &mock.LogFilter{}
		mfmt := &mock.LogFormatter{}
		mw := &mock.Writer{}
		logger := t.setUp(t.message, mf, mfmt, mw)

		logger.Send(t.message)

		mf.AssertExpectations(s.T())
		mfmt.AssertExpectations(s.T())
		mw.AssertExpectations(s.T())
	}
}

func (s *LoggerTestSuite) TestSendNilMessage() {
	mf := &mock.LogFilter{}
	mfmt := &mock.LogFormatter{}
	mw := &mock.Writer{}
	logger := simple.New(mf, mfmt, mw)

	assert.Nil(s.T(), logger.Send(nil))
}

func (s *LoggerTestSuite) TestSendNilFormatter() {
	msg := gol.NewDebug()
	mf := &mock.LogFilter{}
	mf.Mock.On("Filter", msg).Return(false, nil)

	logger := simple.New(mf, nil, nil)

	assert.Error(s.T(), logger.Send(msg))
}

func (s *LoggerTestSuite) TestSendFormatError() {
	msg := gol.NewDebug()
	mf := &mock.LogFilter{}
	mf.Mock.On("Filter", msg).Return(false, nil)
	mfmt := &mock.LogFormatter{}
	mfmt.Mock.On("Format", msg).Return("", fmt.Errorf("unknown"))

	logger := simple.New(mf, mfmt, nil)

	assert.Error(s.T(), logger.Send(msg))
}

func (s *LoggerTestSuite) TestSendNilWriter() {
	msg := gol.NewDebug()
	mf := &mock.LogFilter{}
	mf.Mock.On("Filter", msg).Return(false, nil)
	mfmt := &mock.LogFormatter{}
	mfmt.Mock.On("Format", msg).Return("ERROR", nil)

	logger := simple.New(mf, mfmt, nil)

	assert.Error(s.T(), logger.Send(msg))
}
