package memlib

import (
	"encoding/binary"
	"unsafe"
)

func (p *Process) WriteUint8(address uintptr, value uint8) error {
	return p.WriteMemory(address, []byte{value})
}

func (p *Process) WriteUint16(address uintptr, value uint16) error {
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, value)
	return p.WriteMemory(address, data)
}

func (p *Process) WriteUint32(address uintptr, value uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, value)
	return p.WriteMemory(address, data)
}

func (p *Process) WriteUint64(address uintptr, value uint64) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, value)
	return p.WriteMemory(address, data)
}

func (p *Process) WriteFloat32(address uintptr, value float32) error {
	return p.WriteUint32(address, *(*uint32)(unsafe.Pointer(&value)))
}

func (p *Process) WriteFloat64(address uintptr, value float64) error {
	return p.WriteUint64(address, *(*uint64)(unsafe.Pointer(&value)))
}
