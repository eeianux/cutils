package compress

import (
	"bytes"
	"compress/gzip"
	"io"
)

type gZipImpl struct {
}

func NewGzipImpl() CompressAble {
	return &gZipImpl{}
}

func (g gZipImpl) Compress() (io.WriteCloser, *bytes.Buffer) {
	b := &bytes.Buffer{}
	writer, _ := gzip.NewWriterLevel(b, gzip.BestCompression)
	return writer, b
}

func (g gZipImpl) UnCompress(origin []byte) io.Reader {
	reader, _ := gzip.NewReader(bytes.NewBuffer(origin))
	return reader
}
