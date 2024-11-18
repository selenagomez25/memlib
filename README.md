# memlib

memlib is a cross-platform memory manipulation library written in golang. provides an abstraction layer for dealing with raw memory buffers, specifically designed for reading from and writing to the memory of other processes. this library is useful for game hacking, debugging, and other low-level memory operations.

## Features

- **Process Memory Reading**: Read various data types (e.g., uint8, uint32, float32) from the memory of a specified process.
- **Process Memory Writing**: Write various data types back to the memory of a specified process.
- **Memory Allocation**: Allocate and free memory within the target process.
- **Memory Protection**: Change the protection level of a specific memory region.
- **Module Listing**: List loaded modules within the target process.

## Installation
```bash
go get github.com/selenagomez25/memlib/pkg/memlib
```s