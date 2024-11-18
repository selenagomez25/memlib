package memory

import "golang.org/x/sys/windows"

type MemoryProtector interface {
	Protect(address uintptr, size uint, newProtect uint32) error
}

type BaseProtector struct{}

func (p *BaseProtector) Protect(address uintptr, size uint, newProtect uint32) error {
	var oldProtect uint32
	err := windows.VirtualProtect(address, uintptr(size), newProtect, &oldProtect)
	return err
}
