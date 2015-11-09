# Severity filter

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://raw.githubusercontent.com/mediaFORGE/gol/master/LICENSE)

This example demonstrates the usage of the generic severity filter.

In this case all log messages with severity level error or higher will be written to `stderr`.

```
$ go build example.go
$ go run example.go
severity=EMERGENCY message=system is down
severity=ALERT message=failed to write to disk
severity=CRITICAL message=high server load
severity=ERROR message=invalid number format
```

NOTE: if a message doesn't have a severity level it will also be filtered.
