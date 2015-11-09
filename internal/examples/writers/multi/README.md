# Multi-writer

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://raw.githubusercontent.com/mediaFORGE/gol/master/LICENSE)

This example demonstrates a case where
you would have 1 logger that writes to multiple writers: `stdout` and a file.

```
$ go build example.go
$ go run example.go
message=example execution started
message=example execution ended

$ cat info.log 
message=example execution started
message=example execution ended
```
