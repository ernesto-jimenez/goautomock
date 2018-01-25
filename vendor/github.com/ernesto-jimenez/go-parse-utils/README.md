# go-parse-utils

`go-parse-utils` is a collection of utilities for parsing code easily.

[![GoDoc](https://godoc.org/gopkg.in/src-d/go-parse-utils.v1?status.svg)](https://godoc.org/gopkg.in/src-d/go-parse-utils.v1) [![Build Status](https://travis-ci.org/src-d/go-parse-utils.svg?branch=master)](https://travis-ci.org/src-d/go-parse-utils) [![codecov](https://codecov.io/gh/src-d/go-parse-utils/branch/master/graph/badge.svg)](https://codecov.io/gh/src-d/go-parse-utils) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org) [![Go Report Card](https://goreportcard.com/badge/gopkg.in/src-d/go-parse-utils.v1)](https://goreportcard.com/report/gopkg.in/src-d/go-parse-utils.v1)

### Install

```
go get gopkg.in/src-d/go-parse-utils.v1
```

### Package AST

`PackageAST` retrieves the `*ast.Package` of a package in the given path.

```go
pkg, err := parseutil.PackageAST("github.com/my/project")
```

### Source code importer

The default `importer.Importer` of the Go standard library scans compiled objects, which can be painful to deal with in code generation, as it requires to `go build` before running `go generate`.

This packages provides an implementation of `importer.Importer` and `importer.ImporterFrom` that reads directly from source code if the package is in the GOPATH, otherwise (the stdlib, for example) falls back to the default importer in the standard library.

Features:
* It is safe for concurrent use.
* Caches packages after they are first imported.

```go
importer := parseutil.NewImporter()
pkg, err := importer.Import("github.com/my/project")
```

### License

MIT, see [LICENSE](/LICENSE)
