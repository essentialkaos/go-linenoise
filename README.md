<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-linenoise.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/go-linenoise"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev"></a>
  <a href="https://kaos.sh/w/aligo/ci"><img src="https://kaos.sh/w/aligo/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/aligo"><img src="https://kaos.sh/r/aligo.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/aligo"><img src="https://codebeat.co/badges/f7800a13-657f-4be9-a359-2845f3433588" alt="codebeat badge" /></a>
  <a href="https://kaos.sh/w/aligo/codeql"><img src="https://kaos.sh/w/aligo/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/bsd.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#example">Example</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`go-linenoise` is a Go package wrapping the [linenoise](https://github.com/antirez/linenoise) C library. Since `v3` we use [@yhirose](https://github.com/yhirose) [fork](https://github.com/yhirose/linenoise/tree/utf8-support) with UTF-8 support.

This is fork of [go.linenoise](https://github.com/GeertJohan/go.linenoise) package used in [EK](https://github.com/essentialkaos) projects.

### Installation

For install, do:

```
go get github.com/essentialkaos/go-linenoise
```

For update to latest stable release, do:

```
go get -u github.com/essentialkaos/go-linenoise
```

### Example

```go
package main

// ////////////////////////////////////////////////////////////////////////// //

import (
  "fmt"

  "github.com/essentialkaos/go-linenoise"
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
| `master` | [![CI](https://kaos.sh/w/go-linenoise/ci.svg?branch=master)](https://kaos.sh/w/go-linenoise/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/go-linenoise/ci.svg?branch=develop)](https://kaos.sh/w/go-linenoise/ci?query=branch:develop) |

### License
All code in this repository is licensed under a BSD license.
This project wraps [linenoise](https://github.com/antirez/linenoise) which is written by Salvatore Sanfilippo and Pieter Noordhuis. The license for linenoise is included in the files `linenoise.c` and `linenoise.h`.
For all other files please read the [LICENSE](LICENSE) file.

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
