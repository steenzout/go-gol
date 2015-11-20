![Logo](https://raw.githubusercontent.com/mediaFORGE/gol/develop/logo.png)

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://raw.githubusercontent.com/mediaFORGE/gol/master/LICENSE)
[![Build Status](https://travis-ci.org/mediaFORGE/gol.svg?branch=develop)](https://travis-ci.org/mediaFORGE/gol)

[![Coverage Status](https://coveralls.io/repos/steenzout/gol/badge.svg?branch=develop&service=github)](https://coveralls.io/github/steenzout/gol?branch=develop)

gol is an easily extensible concurrent logging library.

If you would like to contribute
check [CONTRIBUTING.md](https://github.com/mediaFORGE/gol/tree/master/CONTRIBUTING.md).


## Example

```
package main

import (
	"fmt"
	"os"

	"github.com/mediaFORGE/gol"
	"github.com/mediaFORGE/gol/formatters"

	field_severity "github.com/mediaFORGE/gol/fields/severity"
	filter_severity "github.com/mediaFORGE/gol/filters/severity"
	logger_simple "github.com/mediaFORGE/gol/loggers/simple"
	manager_simple "github.com/mediaFORGE/gol/managers/simple"
)

// LogWorkers the number of log message workers.
const LogWorkers = 4

// Log holds the application LogManager instance.
var Log gol.LoggerManager

func init() {
	fmt.Println("init():start")
	Log = manager_simple.New(LogWorkers)

	f := filter_severity.New(field_severity.Info)
	formatter := formatters.Text{}
	logger := logger_simple.New(f, formatter, os.Stdout)
	Log.Register("main", logger)

	Log.Run()
	Log.Send(gol.NewInfo("message", "main.Log has been configured"))
	fmt.Println("init():end")
}

func main() {
	fmt.Println("Started application.")
	defer func() {
		Log.Close()
		fmt.Println("Ended application.")
	}()

	// send 10,000 messages
	for i := 0; i < 100000; i++ {
		Log.Send(gol.NewInfo("i", fmt.Sprintf("%d", i)))
	}
	fmt.Println("Ending application...")
}
```

More examples can be found [here](https://github.com/mediaFORGE/gol/tree/master/internal/examples).
