# Application

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://raw.githubusercontent.com/mediaFORGE/gol/master/LICENSE)

This example demonstrates a case where
you use the log manager to setup your application logging.

```
$ go build .
$ ./application
init():start
init():end
Started application.
severity=INFO i=0
(...)
severity=INFO i=99996
severity=INFO i=99997
severity=INFO i=99998
Ending application...
severity=INFO i=99999
Ended application.
```
