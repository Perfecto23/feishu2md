# lark_rate_limiter

wrap x/time/rate for github.com/chyroc/lark


## Installation

```shell
go get github.com/chyroc/lark_rate_limiter
```

## Usage

```go
package main

import (
	"github.com/chyroc/lark"
	"github.com/chyroc/lark_rate_limiter"
)

// limit lark api request 5 qps
func main() {
	lark.New(
		lark.WithAppCredential("", ""),
		lark.WithApiMiddleware(lark_rate_limiter.Wait(5, 5)),
	)
}
```
