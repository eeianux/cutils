package compress

import (
	"bytes"
	"github.com/golang/snappy"
	"io"
)

type snappyImpl struct {
}

func NewSnappyImpl() CompressAble {
	return &snappyImpl{}
}

func (s snappyImpl) Compress() (io.WriteCloser, *bytes.Buffer) {
	b := &bytes.Buffer{}
	writer := snappy.NewBufferedWriter(b)
	return writer, b
}

func (s snappyImpl) UnCompress(content []byte) io.Reader {
	return snappy.NewReader(bytes.NewBuffer(content))
}
