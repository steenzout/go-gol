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
	"fmt"
	"os"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/formatters"
)

var log gol.Logger

func init() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("ERROR %s\n", err))
	}

	txtFmt := &formatters.Text{}
	log = gol.SimpleLog(nil, txtFmt, file)
}

func main() {
	log.Send(&gol.LogMessage{"message": "example execution started"})
	log.Send(&gol.LogMessage{"message": "example execution ended"})
}
