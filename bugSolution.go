```go
package main

import (

    "fmt"
)

func main() {
    var m map[string]int
    //check if map is nil before access
if m == nil {
        m = make(map[string]int)
    }
    m["key"] = 10 
    fmt.Println(m["key"])

    //comma ok idiom
    val, ok := m["key2"]
    if ok {
        fmt.Println("key2 found", val)
    } else {
        fmt.Println("key2 not found or map is nil")
    }
}
```