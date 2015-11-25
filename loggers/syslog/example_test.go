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
	"log/syslog"
	"net"
	"os"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/formatters"
	mfsyslog "github.com/mediaFORGE/gol/loggers/syslog"
)

var log gol.Logger

func init() {
	log = mfsyslog.New("udp", "127.0.0.1:10001", syslog.LOG_EMERG, "example.go", &formatters.Text{})

	syncch := make(chan bool, 1)
	readch := make(chan bool, 1)
	go ListenToUDP(syncch, readch)
	<-syncch
}

func ListenToUDP(ch chan bool, readch chan bool) {
	if addr, err := net.ResolveUDPAddr("udp", ":10001"); err != nil {
		os.Stderr.WriteString("error ", err)
	} else {
		/* Now listen at selected port */
		if conn, err := net.ListenUDP("udp", addr); err != nil {
			os.Stderr.WriteString("error ", err)
		} else {
			defer conn.Close()

			buf := make([]byte, 1024)

			ch <- true

			for {
				if !<-readch {
					break // end loop execution and end go routine
				}
				msgchan <- m.read(conn, buf)
			}
		}
	}
}

func Example() {
	log.Send(gol.NewInfo("message", "example execution started"))
	// Output: message='example execution started'

	log.Send(gol.NewInfo("message", "example execution ended"))
	// Output: message='example execution ended'
}
