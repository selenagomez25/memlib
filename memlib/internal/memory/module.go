package memory

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type Module struct {
	Name string
	Base uintptr
	Size uint
}

type ModuleLister interface {
	ListModules() ([]Module, error)
}

type BaseModuleLister struct{}

func (m *BaseModuleLister) ListModules() ([]Module, error) {
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPMODULE, 0)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(snapshot)

	var me windows.ModuleEntry32
	me.Size = uint32(unsafe.Sizeof(me))
	err = windows.Module32First(snapshot, &me)
	if err != nil {
		return nil, err
	}

	var modules []Module
	for {
		module := Module{
			Name: windows.UTF16PtrToString(&me.Module[0]),
			Base: uintptr(me.ModBaseAddr),
			Size: uint(me.ModBaseSize),
		}
		modules = append(modules, module)

		err = windows.Module32Next(snapshot, &me)
		if err == windows.ERROR_NO_MORE_FILES {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return modules, nil
}
