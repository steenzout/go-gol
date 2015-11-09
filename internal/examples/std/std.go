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
	"github.com/mediaFORGE/gol/filters"
)

var log gol.Logger = gol.SimpleLog(nil, nil, os.Stdout)
var errorLog gol.Logger = gol.SimpleLog(filters.NewSeverity(severity.Error), nil, os.Stderr)

func main() {
	log.Send(&gol.LogMessage{"hello": "world!"})

	errorLog.Send(gol.NewError("field", "value"))
	errorLog.Send(gol.NewDebug("field_debug", "value"))
}
