package memory

import (
	"encoding/binary"
	"unsafe"

	"golang.org/x/sys/windows"
)

type MemoryWriter interface {
	Write(address uintptr, buffer []byte) (int, error)
	WriteUint8(address uintptr, value uint8) error
	WriteUint16(address uintptr, value uint16) error
	WriteUint32(address uintptr, value uint32) error
	WriteUint64(address uintptr, value uint64) error
	WriteFloat32(address uintptr, value float32) error
	WriteFloat64(address uintptr, value float64) error
}

type BaseWriter struct{}

func (w *BaseWriter) Write(address uintptr, buffer []byte) (int, error) {
	var bytesWritten uintptr
	err := windows.WriteProcessMemory(windows.CurrentProcess(), address, &buffer[0], uintptr(len(buffer)), &bytesWritten)
	return int(bytesWritten), err
}

func (w *BaseWriter) WriteUint8(address uintptr, value uint8) error {
	buffer := []byte{value}
	_, err := w.Write(address, buffer)
	return err
}

func (w *BaseWriter) WriteUint16(address uintptr, value uint16) error {
	buffer := make([]byte, 2)
	binary.LittleEndian.PutUint16(buffer, value)
	_, err := w.Write(address, buffer)
	return err
}

func (w *BaseWriter) WriteUint32(address uintptr, value uint32) error {
	buffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, value)
	_, err := w.Write(address, buffer)
	return err
}

func (w *BaseWriter) WriteUint64(address uintptr, value uint64) error {
	buffer := make([]byte, 8)
	binary.LittleEndian.PutUint64(buffer, value)
	_, err := w.Write(address, buffer)
	return err
}

func (w *BaseWriter) WriteFloat32(address uintptr, value float32) error {
	return w.WriteUint32(address, *(*uint32)(unsafe.Pointer(&value)))
}

func (w *BaseWriter) WriteFloat64(address uintptr, value float64) error {
	return w.WriteUint64(address, *(*uint64)(unsafe.Pointer(&value)))
}
