package memory

import (
	"encoding/binary"
	"unsafe"

	"golang.org/x/sys/windows"
)

type MemoryReader interface {
	Read(address uintptr, buffer []byte) (int, error)
	ReadUint8(address uintptr) (uint8, error)
	ReadUint16(address uintptr) (uint16, error)
	ReadUint32(address uintptr) (uint32, error)
	ReadUint64(address uintptr) (uint64, error)
	ReadFloat32(address uintptr) (float32, error)
	ReadFloat64(address uintptr) (float64, error)
}

type BaseReader struct{}

func (r *BaseReader) Read(address uintptr, buffer []byte) (int, error) {
	var bytesRead uintptr
	err := windows.ReadProcessMemory(windows.CurrentProcess(), address, &buffer[0], uintptr(len(buffer)), &bytesRead)
	return int(bytesRead), err
}

func (r *BaseReader) ReadUint8(address uintptr) (uint8, error) {
	var buffer [1]byte
	_, err := r.Read(address, buffer[:])
	if err != nil {
		return 0, err
	}
	return buffer[0], nil
}

func (r *BaseReader) ReadUint16(address uintptr) (uint16, error) {
	var buffer [2]byte
	_, err := r.Read(address, buffer[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(buffer[:]), nil
}

func (r *BaseReader) ReadUint32(address uintptr) (uint32, error) {
	var buffer [4]byte
	_, err := r.Read(address, buffer[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(buffer[:]), nil
}

func (r *BaseReader) ReadUint64(address uintptr) (uint64, error) {
	var buffer [8]byte
	_, err := r.Read(address, buffer[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(buffer[:]), nil
}

func (r *BaseReader) ReadFloat32(address uintptr) (float32, error) {
	value, err := r.ReadUint32(address)
	if err != nil {
		return 0, err
	}
	return *(*float32)(unsafe.Pointer(&value)), nil
}

func (r *BaseReader) ReadFloat64(address uintptr) (float64, error) {
	value, err := r.ReadUint64(address)
	if err != nil {
		return 0, err
	}
	return *(*float64)(unsafe.Pointer(&value)), nil
}
