package fileExample

import (
	"bytes"
	"encoding/binary"
	"os"
)

type payload struct {
	One   float32
	Two   float64
	Three uint32
}

type WriteFileDemo struct{}

func NewWriteFileDemo() *WriteFileDemo {
	return &WriteFileDemo{}
}

func (w *WriteFileDemo) WriteFile(file *os.File, binInfo []byte) error {
	_, err := file.Write(binInfo)
	return err
}

func (w *WriteFileDemo) WriteBinaryFile(dir string) error {
	createBinFile, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer createBinFile.Close()
	for i := 0; i < 50; i++ {
		s := &payload{
			1.10,
			2.20,
			3,
		}
		var binBuf bytes.Buffer
		binary.Write(&binBuf, binary.BigEndian, s)
		w.WriteFile(createBinFile, binBuf.Bytes())
	}
	return nil
}
