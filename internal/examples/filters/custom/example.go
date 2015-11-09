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

package main

import (
	"os"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/fields/severity"
	"github.com/mediaFORGE/gol/formatters"
)

type ErrorFilter struct{}

func (f *ErrorFilter) Filter(msg *gol.LogMessage) bool {
	if s, err := msg.Severity(); err != nil {
		// no severity
		return true
	} else {
		return s >= severity.Warning
	}
}

var _ gol.LogFilter = (*ErrorFilter)(nil)

type InfoFilter struct{}

func (f *InfoFilter) Filter(msg *gol.LogMessage) bool {
	if s, err := msg.Severity(); err != nil {
		// no severity
		return true
	} else {
		return s < severity.Warning
	}
}

var _ gol.LogFilter = (*InfoFilter)(nil)

var txtFmt = &formatters.Text{}
var log gol.Logger = gol.SimpleLog(&InfoFilter{}, txtFmt, os.Stdout)
var errorLog gol.Logger = gol.SimpleLog(&ErrorFilter{}, txtFmt, os.Stderr)

func main() {
	log.Send(&gol.LogMessage{"hello": "world!"})
	log.Send(gol.NewError("error", "info"))
	log.Send(gol.NewDebug("debug", "info"))

	errorLog.Send(&gol.LogMessage{"hello": "world!"})
	errorLog.Send(gol.NewEmergency("emergency", "errorLog"))
	errorLog.Send(gol.NewError("error", "errorLog"))
	errorLog.Send(gol.NewWarning("error", "errorLog"))
	errorLog.Send(gol.NewInfo("info", "errorLog"))
	errorLog.Send(gol.NewDebug("debug", "errorLog"))
}
