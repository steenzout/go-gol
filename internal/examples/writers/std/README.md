# Standard

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://raw.githubusercontent.com/mediaFORGE/gol/master/LICENSE)

This example demonstrates a case where
you would have 2 loggers implemented that write to `stdout` and `stderr`, respectively.

```
$ go run example.go > info.log 2> error.log

$ ls 
README.md		example		example.go	error.log			info.log

$ cat info.log 
message=written to log

$ cat err.log 
message=written to error log
```
