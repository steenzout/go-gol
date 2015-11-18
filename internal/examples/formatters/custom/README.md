# Custom formatter

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://raw.githubusercontent.com/mediaFORGE/gol/master/LICENSE)

This example demonstrates a case where
you have a logger that writes to `stdout` using a custom message formatter.

```
# install dependencies
$ go get github.com/fatih/color

# build
$ go buid example.go

# run
$ ./custom
severity:EMERGENCY message:system is down
severity:ALERT message:failed to write to disk
severity:CRITICAL message:high server load
severity:ERROR message:invalid number format
message:performance close to 1s threshold severity:WARNING
severity:NOTICE message:failed to communicate with monitoring service
severity:INFO message:requested processed in 250ms
severity:DEBUG debug:var x = 10
```
