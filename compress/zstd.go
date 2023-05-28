package compress

import (
	"bytes"
	"github.com/klauspost/compress/zstd"
	"io"
)

type zstdImpl struct {
}

func NewZstdImpl() CompressAble {
	return &zstdImpl{}
}

func (z zstdImpl) Compress() (io.WriteCloser, *bytes.Buffer) {
	b := &bytes.Buffer{}
	writer, _ := zstd.NewWriter(b, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
	return writer, b
}

func (z zstdImpl) UnCompress(origin []byte) io.Reader {
	reader, _ := zstd.NewReader(bytes.NewBuffer(origin))
	return reader
}
