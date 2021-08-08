## cuid

Collision-resistant hashes for the cloud, in Go.

The `cuid` package provides collision-resistant ids optimized for horizontal scaling and sequential lookup performance. This README file is just going to cover the basics and the Go-specific implementation details.

**Please refer to the [main project site](http://usecuid.org) for the full rationale behind CUIDs.**

## Sample CUID

ch72gsb320000udocl363eofy

## Sample CUID Slug

ew0k9fwpl

# Installation

`go get -u github.com/lucsky/cuid`

## Example usage

```Go
package main

import (
    "crypto/rand"
    "fmt"

    "github.com/lucsky/cuid"
)

func main() {
    // Generate pseudo-random CUID
    fmt.Println(cuid.New())
    // Generate slug
    fmt.Println(cuid.Slug())

    // Generate cryptographic-random CUID
    c, err := cuid.NewCrypto(rand.Reader)
    if err != nil {
        fmt.Printf("%v", err)
        return
    }
    fmt.Println(c)
}
```

## Go package specific features

The Go cuid package provides APIs to specify a custom random source as well as a custom counter. A custom counter implementation could provide a centralized Redis base counter, for example.

## Contributors

-   Luc Heinrich (@lucsky, author)
-   Thomas Hopkins (@hopkinsth)
-   fiatjaf
-   Marcus Dantas (@mpsdantas)
-   Oscar Forner Martinez (@maitesin)
-   Mike Frey (@mikefrey)
-   Andris Mednis (@andrismednis)
