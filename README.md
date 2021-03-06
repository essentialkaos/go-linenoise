<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-linenoise.svg"/></a></p>

<p align="center">
  <a href="https://pkg.go.dev/pkg.re/essentialkaos/go-linenoise.v3"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/go-linenoise"><img src="https://goreportcard.com/badge/github.com/essentialkaos/go-linenoise"></a>
  <a href="https://github.com/essentialkaos/go-linenoise/actions"><img src="https://github.com/essentialkaos/go-linenoise/workflows/CI/badge.svg" alt="GitHub Actions Status" /></a>
  <a href="https://github.com/essentialkaos/go-linenoise/actions?query=workflow%3ACodeQL"><img src="https://github.com/essentialkaos/go-linenoise/workflows/CodeQL/badge.svg" /></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-go-linenoise"><img alt="codebeat badge" src="https://codebeat.co/badges/f7800a13-657f-4be9-a359-2845f3433588" /></a>
  <a href="https://github.com/essentialkaos/go-linenoise/blob/master/LICENSE"><img src="https://gh.kaos.st/bsd.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#example">Example</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`go-linenoise` is a Go package wrapping the [linenoise](https://github.com/antirez/linenoise) C library. Since `v3` we use [@yhirose](https://github.com/yhirose) [fork](https://github.com/yhirose/linenoise/tree/utf8-support) with UTF-8 support.

This is fork of [go.linenoise](https://github.com/GeertJohan/go.linenoise) package used in [EK](https://github.com/essentialkaos) projects.

### Installation

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
| `master` | [![CI](https://github.com/essentialkaos/go-linenoise/workflows/CI/badge.svg?branch=master)](https://github.com/essentialkaos/go-linenoise/actions) |
| `develop` | [![CI](https://github.com/essentialkaos/go-linenoise/workflows/CI/badge.svg?branch=develop)](https://github.com/essentialkaos/go-linenoise/actions) |

### License
All code in this repository is licensed under a BSD license.
This project wraps [linenoise](https://github.com/antirez/linenoise) which is written by Salvatore Sanfilippo and Pieter Noordhuis. The license for linenoise is included in the files `linenoise.c` and `linenoise.h`.
For all other files please read the [LICENSE](LICENSE) file.

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
