//go:build windows
// +build windows

package memlib

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Module struct {
	Name string
	Base uintptr
	Size uint
}

func (p *Process) ListModules() ([]Module, error) {
	var modules []Module
	var mod windows.ModuleEntry32
	mod.Size = uint32(unsafe.Sizeof(mod))
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPMODULE|windows.TH32CS_SNAPMODULE32, uint32(p.pid))
	if err != nil {
		return nil, fmt.Errorf("CreateToolhelp32Snapshot failed: %v", err)
	}
	defer windows.CloseHandle(snapshot)

	err = windows.Module32First(snapshot, &mod)
	if err != nil {
		return nil, fmt.Errorf("Module32First failed: %v", err)
	}

	for {
		modules = append(modules, Module{
			Name: windows.UTF16PtrToString(&mod.Module[0]),
			Base: uintptr(mod.ModBaseAddr),
			Size: uint(mod.ModBaseSize),
		})

		err = windows.Module32Next(snapshot, &mod)
		if err == windows.ERROR_NO_MORE_FILES {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Module32Next failed: %v", err)
		}
	}

	return modules, nil
}
