package compress

import (
	"bytes"
	"compress/zlib"
	"io"
)

type zLibImpl struct {
}

func NewZLibImpl() CompressAble {
	return &zLibImpl{}
}

func (z zLibImpl) Compress() (io.WriteCloser, *bytes.Buffer) {
	b := &bytes.Buffer{}
	writer, _ := zlib.NewWriterLevel(b, zlib.BestCompression)
	return writer, b
}

func (z zLibImpl) UnCompress(content []byte) io.Reader {
	reader, _ := zlib.NewReader(bytes.NewBuffer(content))
	return reader
}
