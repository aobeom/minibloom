# minibloom

Simple bloom filter using murmurhash3.

## Example

```go
package main

import (
    "fmt"

    "github.com/aobeom/minibloom"
)

func main() {
    size := 2 << 20
    counts := 7
    bloom := minibloom.New(size, counts)

    key := "test1"

    bloom.Add([]byte(key))

    if bloom.In([]byte(key)) {
        fmt.Println("exist")
    }
}
```
