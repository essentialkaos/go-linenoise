# `go-linenoise` [![GoDoc](https://godoc.org/pkg.re/essentialkaos/go-linenoise.v2?status.svg)](https://godoc.org/pkg.re/essentialkaos/go-linenoise.v2) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/go-linenoise)](https://goreportcard.com/report/github.com/essentialkaos/go-linenoise)

go-linenoise is a [go](http://golang.org) package wrapping the [linenoise](https://github.com/antirez/linenoise) C library.

This package does not compile on windows.

## Instalation

```
go get pkg.re/essentialkaos/go-linenoise
```

For update to latest stable release, do:

```
go get -u pkg.re/essentialkaos/go-linenoise
```

## License
All code in this repository is licensed under a BSD license.
This project wraps [linenoise](https://github.com/antirez/linenoise) which is written by Salvatore Sanfilippo and Pieter Noordhuis. The license for linenoise is included in the files `linenoise.c` and `linenoise.h`.
For all other files please read the [LICENSE](LICENSE) file.