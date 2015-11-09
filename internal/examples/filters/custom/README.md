# Custom filter

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://raw.githubusercontent.com/mediaFORGE/gol/master/LICENSE)

This example demonstrates a case where you would have 2 loggers implemented with 2 different filters.

The first logger is an error logger and it will write to `stderr` all messages higher than `WARNING`.

The second logger is an info logger and it will write to `stdout` all messages lower or equal to `WARNING`.

```
$ go build example.go
$ go run example.go > info.log 2> error.log

$ ls 
README.md		example		example.go	error.log			info.log

$ cat error.log 
severity=EMERGENCY emergency=errorLog
severity=ERROR error=errorLog

$ cat info.log 
severity=DEBUG debug=info
```

NOTE: if a message doesn't have a severity level it will also be filtered.
