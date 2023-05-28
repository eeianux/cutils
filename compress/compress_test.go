package compress

import (
	"encoding/base64"
	"fmt"
	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func randStr() string {
	s := ""
	for i := 0; i < 10000; i++ {
		s = s + fmt.Sprint(rand.Int())
	}
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func TestCompressAble(t *testing.T) {
	PatchConvey("rand test", t, func() {
		origin := randStr()
		implList := []CompressAble{NewZstdImpl(), NewGzipImpl(), NewZLibImpl(), NewSnappyImpl()}
		var l []int
		for _, c := range implList {
			compressStr := Compress(origin, c)
			//fmt.Println(compressStr)
			So(compressStr, ShouldNotBeEmpty)
			u := UnCompress(compressStr, c)
			//fmt.Println(u)
			So(u, ShouldNotBeEmpty)
			So(origin, ShouldEqual, u)
			l = append(l, len(compressStr))
		}
		fmt.Println(len(origin))
		fmt.Println(l)
	})
}
