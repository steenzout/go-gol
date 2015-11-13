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

package syslog_test

import (
	"fmt"
	"log/syslog"
	"net"
	"strings"
	"testing"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/fields/severity"
	"github.com/mediaFORGE/gol/formatters"
	golmock "github.com/mediaFORGE/gol/internal/mock"
	golsys "github.com/mediaFORGE/gol/loggers/syslog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockUDPServer
type MockUDPServer struct {
	mock.Mock
	t *testing.T
}

func (m *MockUDPServer) ReadMessage(conn *net.UDPConn, buf []byte) {
	n, addr, err := conn.ReadFromUDP(buf)
	assert.Nil(m.t, err)
	assert.NotNil(m.t, "abc", addr)
	assert.True(m.t, n > 0)

	msg := string(buf[0:n])
	m.receivedMessage(msg)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func (m *MockUDPServer) receivedMessage(msg string) {
	m.Mock.Called(msg)
}

func (m *MockUDPServer) Run(ch chan bool, readch chan bool) {
	if addr, err := net.ResolveUDPAddr("udp", ":10001"); err != nil {
		assert.Fail(m.t, "could not resolve UDP server address")
	} else {
		/* Now listen at selected port */
		if conn, err := net.ListenUDP("udp", addr); err != nil {
			assert.Fail(m.t, "could not listen to UDP server port")
		} else {
			defer conn.Close()

			buf := make([]byte, 1024)

			ch <- true

			for {
				if !<-readch {
					break // end loop execution and end go routine
				}
				m.ReadMessage(conn, buf)
			}
		}
	}
}

// SyslogTestSuite test suite for the github.com/mediaFORGE/gol/loggers/syslog package.
type SyslogTestSuite struct {
	suite.Suite
	logger gol.Logger
}

func (s *SyslogTestSuite) SetupTest() {
	s.logger = golsys.New("udp", ":10001", syslog.LOG_INFO, "test", &formatters.Text{})
}

func (s *SyslogTestSuite) TestSend() {
	udpserver := MockUDPServer{t: s.T()}
	udpserver.Mock.On("receivedMessage", mock.Anything).Return(nil)

	syncch := make(chan bool, 1)
	readch := make(chan bool, 1)
	go udpserver.Run(syncch, readch)
	<-syncch

	var lvl severity.Type
	for lvl = severity.Emergency; lvl <= severity.Debug; lvl++ {
		s.logger.Send(gol.NewMessage(lvl, "message", "unknown"))
		readch <- true // read message sent
	}

	// invalid severity will be sent with the default logger severity level
	s.logger.Send(gol.NewMessage(severity.Debug+1, "invalid", "severity"))
	readch <- true  // read message sent
	readch <- false // end concurrent go routine

	udpserver.AssertExpectations(s.T())

	found := 0
	for lvl = severity.Emergency; lvl <= severity.Debug; lvl++ {
		for _, call := range udpserver.Calls {
			if call.Method == "receivedMessage" {
				for _, arg := range call.Arguments {
					if strings.Contains(arg.(string), lvl.String()) {
						found++
					}
				}
			}
		}
	}
	assert.Equal(s.T(), 8, found)
}

func (s *SyslogTestSuite) TestSendMessageWithoutSeverity() {
	assert.Error(s.T(), s.logger.Send(&gol.LogMessage{"message": "unknown"}))
}

func (s *SyslogTestSuite) TestSendNil() {
	assert.Nil(s.T(), s.logger.Send(nil))
}

func (s *SyslogTestSuite) TestSendNilFormatter() {
	// reset logger built by SetupTest
	s.logger = golsys.New("udp", ":10001", syslog.LOG_INFO, "test", nil)
	assert.Error(s.T(), s.logger.Send(gol.NewEmergency("message", "unknown")))
}

func (s *SyslogTestSuite) TestSendFormatterError() {
	msg := gol.NewEmergency("message", "unknown")
	m := &golmock.MockLogFormatter{}
	m.On("Format", msg).Return("", fmt.Errorf("internal error"))

	s.logger.SetFormatter(m)
	assert.Error(s.T(), s.logger.Send(msg))

	m.AssertExpectations(s.T())
}
