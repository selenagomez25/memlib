package memory

import (
	"golang.org/x/sys/windows"
)

type MemoryAllocator interface {
	Allocate(size uint) (uintptr, error)
	Free(address uintptr) error
}

type BaseAllocator struct{}

func (a *BaseAllocator) Allocate(size uint) (uintptr, error) {
	addr, err := windows.VirtualAlloc(0, uintptr(size), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if err != nil {
		return 0, err
	}
	return addr, nil
}

func (a *BaseAllocator) Free(address uintptr) error {
	return windows.VirtualFree(address, 0, windows.MEM_RELEASE)
}
