<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-linenoise.svg"/></a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/go-linenoise.v3"><img src="https://godoc.org/pkg.re/essentialkaos/go-linenoise.v3?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/go-linenoise"><img src="https://goreportcard.com/badge/github.com/essentialkaos/go-linenoise"></a>
  <a href="https://travis-ci.org/essentialkaos/go-linenoise"><img src="https://travis-ci.org/essentialkaos/go-linenoise.svg"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-go-linenoise"><img alt="codebeat badge" src="https://codebeat.co/badges/f7800a13-657f-4be9-a359-2845f3433588" /></a>
  <a href="https://github.com/essentialkaos/go-linenoise/blob/master/LICENSE"><img src="https://gh.kaos.st/bsd.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#example">Example</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`go-linenoise` is a Go package wrapping the [linenoise](https://github.com/antirez/linenoise) C library. Since `v3` we use [@yhirose](https://github.com/yhirose) [fork](https://github.com/yhirose/linenoise/tree/utf8-support) with UTF-8 support.

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

```go
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
    return
  }

  fmt.Printf("Input: %s\n", input)
}

```

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
