# go-parallel

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mandykoh/go-parallel)](https://pkg.go.dev/github.com/mandykoh/go-parallel)
[![Go Report Card](https://goreportcard.com/badge/github.com/mandykoh/go-parallel)](https://goreportcard.com/report/github.com/mandykoh/go-parallel)
[![Build Status](https://travis-ci.org/mandykoh/go-parallel.svg?branch=main)](https://travis-ci.org/mandykoh/go-parallel)

Simple parallel processing utilities for Go.

See the [API documentation](https://pkg.go.dev/github.com/mandykoh/go-parallel) for more details.

This software is made available under an [MIT license](LICENSE).


## Example usage

Split processing of a large number of things across eight workers:

```go
var thingsToProcess []Thing

parallel.RunWorkers(8, func(workerNum, workerCount int) {
    for i := workerNum; i < len(thingsToProcess); i += workerCount {
        processThing(thingsToProcess[i])
    }
})
```

`RunWorkers` returns when all workers have run to completion, allowing a task to be performed in parallel but treated more simply as a synchronous call.