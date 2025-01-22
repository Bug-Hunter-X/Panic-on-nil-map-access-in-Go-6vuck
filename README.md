# Panic on nil map access in Go

This repository demonstrates a common error in Go: panicking when trying to access a key in a nil map.  The `bug.go` file shows the problematic code. The solution, `bugSolution.go`, demonstrates how to safely handle nil map scenarios.

## The Problem

In Go, attempting to access a key from an uninitialized map (a nil map) will result in a runtime panic. This can be difficult to debug, especially in complex applications.

## The Solution

The solution lies in checking for a nil map before attempting access. Use the following strategies:

1. **Explicit nil check:** Before accessing the map, verify whether it's nil.
2. **Using the comma ok idiom:**  This is a more idiomatic approach.  It safely retrieves the value associated with the key. If the key is not found or the map is nil, the second return value will be false, allowing for graceful handling.

This repository highlights the importance of robust error handling when working with maps in Go, preventing unexpected panics and improving code stability.