## cuid

Collision-resistant hashes for the cloud, in Go.

The ```cuid`` package provides collision-resistant ids optimized for horizontal scaling and sequential lookup performance. This README file is just going to cover the basics and the Go-specific implementation details.

**Please refer to the [main project site](http://usecuid.org) for the full rationale behine CUIDs.**

## Sample CUID

ch72gsb320000udocl363eofy

## Example usage

```Go
package main

import fmt
import "gopkg.in/lucsky/cuid.v1"

func main() {
    fmt.Println(cuid.New())
}
```

## Go package specific features

The Go cuid package provides APIs to specify a custom random source as well as a custom counter source.
