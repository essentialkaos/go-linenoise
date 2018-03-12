# `go-linenoise` [![GoDoc](https://godoc.org/pkg.re/essentialkaos/go-linenoise.v3?status.svg)](https://godoc.org/pkg.re/essentialkaos/go-linenoise.v3) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/go-linenoise)](https://goreportcard.com/report/github.com/essentialkaos/go-linenoise) [![codebeat badge](https://codebeat.co/badges/f7800a13-657f-4be9-a359-2845f3433588)](https://codebeat.co/projects/github-com-essentialkaos-go-linenoise) [![License](https://gh.kaos.st/bsd.svg)](LICENSE)

go-linenoise is a Go package wrapping the [linenoise](https://github.com/antirez/linenoise) C library. Since `v3` we use [@yhirose](https://github.com/yhirose) [fork](https://github.com/yhirose/linenoise/tree/utf8-support) with UTF-8 support.

This is fork of [go.linenoise](https://github.com/GeertJohan/go.linenoise) package used in [EK](https://github.com/essentialkaos) projects.

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

For install, do:

```
go get pkg.re/essentialkaos/go-linenoise.v3
```

For update to latest stable release, do:

```
go get -u pkg.re/essentialkaos/go-linenoise.v3
```

### Example

````go
package main

// ////////////////////////////////////////////////////////////////////////// //

import (
  "fmt"

  "pkg.re/essentialkaos/go-linenoise.v3"
)

// ////////////////////////////////////////////////////////////////////////// //

func main() {
  input, err := linenoise.Line("> ")

  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }

  fmt.Printf("Input: %s\n", input)
}

````

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/go-linenoise.svg?branch=master)](https://travis-ci.org/essentialkaos/go-linenoise) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/go-linenoise.svg?branch=develop)](https://travis-ci.org/essentialkaos/go-linenoise) |

### License
All code in this repository is licensed under a BSD license.
This project wraps [linenoise](https://github.com/antirez/linenoise) which is written by Salvatore Sanfilippo and Pieter Noordhuis. The license for linenoise is included in the files `linenoise.c` and `linenoise.h`.
For all other files please read the [LICENSE](LICENSE) file.

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
