# fileutil

[![GoDoc](https://pkg.go.dev/badge/github.com/KarpelesLab/fileutil)](https://pkg.go.dev/github.com/KarpelesLab/fileutil)

A Go package providing utilities for safe and atomic file operations.

## Installation

```bash
go get github.com/KarpelesLab/fileutil
```

## Features

- **Put**: Conditionally write data only if content differs, with atomic operations
- **WriteFileReader**: Atomically write data from an io.Reader using a temporary file with `~` suffix
- **TarExtract**: Extract tar archives to a directory

## Documentation

Full API documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/KarpelesLab/fileutil).