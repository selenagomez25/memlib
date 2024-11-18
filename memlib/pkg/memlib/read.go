package memlib

import (
	"encoding/binary"
	"unsafe"
)

func (p *Process) ReadUint8(address uintptr) (uint8, error) {
	data, err := p.ReadMemory(address, 1)
	if err != nil {
		return 0, err
	}
	return data[0], nil
}

func (p *Process) ReadUint16(address uintptr) (uint16, error) {
	data, err := p.ReadMemory(address, 2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(data), nil
}

func (p *Process) ReadUint32(address uintptr) (uint32, error) {
	data, err := p.ReadMemory(address, 4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(data), nil
}

func (p *Process) ReadUint64(address uintptr) (uint64, error) {
	data, err := p.ReadMemory(address, 8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(data), nil
}

func (p *Process) ReadFloat32(address uintptr) (float32, error) {
	data, err := p.ReadUint32(address)
	if err != nil {
		return 0, err
	}
	return *(*float32)(unsafe.Pointer(&data)), nil
}

func (p *Process) ReadFloat64(address uintptr) (float64, error) {
	data, err := p.ReadUint64(address)
	if err != nil {
		return 0, err
	}
	return *(*float64)(unsafe.Pointer(&data)), nil
}
