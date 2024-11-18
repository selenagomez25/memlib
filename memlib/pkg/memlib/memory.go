//go:build windows
// +build windows

package memlib

import (
	"fmt"

	"golang.org/x/sys/windows"
)

const PROCESS_ALL_ACCESS = 0x1F0FFF

type Process struct {
	pid    int
	handle windows.Handle
}

func NewProcess(pid int) (*Process, error) {
	handle, err := windows.OpenProcess(PROCESS_ALL_ACCESS, false, uint32(pid))
	if err != nil {
		return nil, fmt.Errorf("OpenProcess failed: %v", err)
	}
	return &Process{pid: pid, handle: handle}, nil
}

func (p *Process) Close() error {
	return windows.CloseHandle(p.handle)
}

func (p *Process) ReadMemory(address uintptr, size uint) ([]byte, error) {
	buffer := make([]byte, size)
	var bytesRead uintptr
	err := windows.ReadProcessMemory(p.handle, address, &buffer[0], uintptr(size), &bytesRead)
	if err != nil {
		return nil, fmt.Errorf("ReadProcessMemory failed: %v", err)
	}
	return buffer[:bytesRead], nil
}

func (p *Process) WriteMemory(address uintptr, data []byte) error {
	var bytesWritten uintptr
	err := windows.WriteProcessMemory(p.handle, address, &data[0], uintptr(len(data)), &bytesWritten)
	if err != nil {
		return fmt.Errorf("WriteProcessMemory failed: %v", err)
	}
	if bytesWritten != uintptr(len(data)) {
		return fmt.Errorf("WriteProcessMemory wrote %d bytes, expected %d", bytesWritten, len(data))
	}
	return nil
}
