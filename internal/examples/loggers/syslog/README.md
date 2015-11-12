# Syslog

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://raw.githubusercontent.com/mediaFORGE/gol/master/LICENSE)

This example demonstrates a case where
you have a logger that writes to syslog.

```
$ go build example.go
$ go run example.go
message= severity=INFO message=example execution started

level=Info
message= severity=INFO message=example execution ended

level=Info
```
