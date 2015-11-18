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

	"github.com/mediaFORGE/gol"
)

func main() {
	fmt.Println("Started application.")
	defer func() {
		Log.Close()
		fmt.Println("Application ended.")
	}()

	// send 2,000 messages
	for i := 0; i < 2000; i++ {
		Log.Send(gol.NewInfo("i", fmt.Sprintf("%d", i)))
	}
	fmt.Println("Ending application...")
}
