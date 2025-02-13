package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

func Ttr() string {
	str := strings.Repeat("A", 10*1024*1024) // Repeat "A" n times
	b := str[100:200]                        // We don't create a new string, but refer to str
	fmt.Println("Created a large variable")
	return b
}

func main() {
	var m runtime.MemStats

	// Run the function
	b := Ttr()

	// Check memory usage before GC
	runtime.ReadMemStats(&m)
	fmt.Printf("Before GC: %v MB\n", m.Alloc/1024/1024)

	// Run garbage collector and free memory in OS
	runtime.GC()
	debug.FreeOSMemory()

	// Check memory usage after GC
	runtime.ReadMemStats(&m)
	fmt.Printf("After GC: %v MB\n", m.Alloc/1024/1024)

	fmt.Println(b)
}
