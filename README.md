# must

![License](https://img.shields.io/badge/license-Apache2.0-green)
![Language](https://img.shields.io/badge/Language-Go-blue.svg)
[![version](https://img.shields.io/github/v/tag/libgox/must?label=release&color=blue)](https://github.com/libgox/must/releases)
[![Godoc](http://img.shields.io/badge/docs-go.dev-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/libgox/must)
[![Go report](https://goreportcard.com/badge/github.com/libgox/must)](https://goreportcard.com/report/github.com/libgox/must)
[![codecov](https://codecov.io/gh/libgox/must/branch/main/graph/badge.svg)](https://codecov.io/gh/libgox/must)

## 📋 Requirements

- Go 1.18+

## 🚀 Install

```
go get github.com/libgox/must
```

## 💡 Usage

Must package provides a set of utility functions that simplify error handling when dealing with functions that returns both a value and an error.
These functions will panic if the error is not `nil`, otherwise, they return the provided value(s).

The `Must` function returns the value when no error is encountered. If an error is encountered, it panics with the provided error.

```go
package main

import (
	"fmt"
	"github.com/libgox/must"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	// Will panic if error is not nil
	must.Must(file, err)
	defer file.Close()

	// Proceed with file operations
}

```

`Must2` and `Must3` are variant for different number of return values
