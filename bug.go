```go
func main() {
    var m map[string]int
    m["key"] = 10 // This will panic if the map is nil
    fmt.Println(m["key"])
}
```