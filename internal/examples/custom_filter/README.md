# custom_filter

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://raw.githubusercontent.com/mediaFORGE/gol/master/LICENSE)

This example demonstrates a case where you would have 2 loggers implemented with different filters and writers.

The first logger is an error logger and it will write to `stderr` all messages higher than `WARNING`.

The second logger is an info logger and it will write to `stdout` all messages lower or equal to `WARNING`.

```
$ go run custom_filter.go > info.log 2> err.log

$ ls 
README.md		custom_filter		custom_filter.go	err.log			info.log

$ cat info.log 
{"debug":"info","severity":7}

$ cat err.log 
emergency=errorLog severity=EMERGENCY
severity=ERROR error=errorLog
```

Note that if a message doesn't have a severity level it will also be filtered.
