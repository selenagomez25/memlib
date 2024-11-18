# memlib

memlib is a cross-platform memory hacking library written in Go. It provides an abstraction layer for dealing with raw memory buffers, especially when dealing with process memory.

## Features

- Cross-platform support
- Memory reading and writing
- Memory protection
- Memory allocation
- Module listing
- Overlay rendering (to be implemented)

## Usage

```go
import "github.com/selenagomez25/memlib/pkg/memlib"

func main() {
    process := memlib.NewProcess(1234) // Replace with actual PID
    
    address := uintptr(0x12345678)
    value, err := process.ReadUint32(address)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Value at 0x%x: %d\n", address, value)
}
```