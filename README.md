# minibloom

Simple bloom filter using murmurhash3.

## Example

```go
package main

import "github.com/aobeom/minibloom"

size := 2 << 20
counts := 7
bloom := minibloom.New(size, counts)

key := "test1"

bloom.Add([]byte(key))

if bloom.In([]byte(key)) {
    fmt.Println("exist")
}
```
