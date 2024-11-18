//go:build windows
// +build windows

package memlib

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func (p *Process) ProtectMemory(address uintptr, size uint, newProtect uint32) error {
	var oldProtect uint32
	err := windows.VirtualProtectEx(p.handle, address, uintptr(size), newProtect, &oldProtect)
	if err != nil {
		return fmt.Errorf("VirtualProtectEx failed: %v", err)
	}
	return nil
}
