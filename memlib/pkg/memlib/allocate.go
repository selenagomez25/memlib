//go:build windows
// +build windows

package memlib

import (
	"fmt"

	"golang.org/x/sys/windows"
)

const (
	MEM_COMMIT  = 0x1000
	MEM_RESERVE = 0x2000
	MEM_RELEASE = 0x8000
)

var (
	modkernel32        = windows.NewLazySystemDLL("kernel32.dll")
	procVirtualAllocEx = modkernel32.NewProc("VirtualAllocEx")
	procVirtualFreeEx  = modkernel32.NewProc("VirtualFreeEx")
)

func (p *Process) AllocateMemory(size uint) (uintptr, error) {
	addr, _, err := procVirtualAllocEx.Call(
		uintptr(p.handle),
		0,
		uintptr(size),
		MEM_COMMIT|MEM_RESERVE,
		windows.PAGE_READWRITE,
	)
	if addr == 0 {
		return 0, fmt.Errorf("VirtualAllocEx failed: %v", err)
	}
	return addr, nil
}

func (p *Process) FreeMemory(address uintptr) error {
	ret, _, err := procVirtualFreeEx.Call(
		uintptr(p.handle),
		address,
		0,
		MEM_RELEASE,
	)
	if ret == 0 {
		return fmt.Errorf("VirtualFreeEx failed: %v", err)
	}
	return nil
}
