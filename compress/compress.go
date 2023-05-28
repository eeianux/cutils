package compress

import (
	"bytes"
	"encoding/base64"
	"io"
)

type CompressAble interface {
	Compress() (io.WriteCloser, *bytes.Buffer)
	UnCompress(content []byte) io.Reader
}

func Compress(origin string, able CompressAble) string {
	writer, b := able.Compress()
	_, _ = writer.Write([]byte(origin))
	_ = writer.Close()
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func UnCompress(origin string, able CompressAble) string {
	c, _ := base64.StdEncoding.DecodeString(origin)
	reader := able.UnCompress(c)
	b, _ := io.ReadAll(reader)
	return string(b)
}
