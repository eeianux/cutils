package test

import (
	"fmt"
	"io"
	"os"
	"testing"

	. "github.com/bytedance/mockey"
	"github.com/eeianux/cutils/excel"
	. "github.com/smartystreets/goconvey/convey"
)

func compare2File(a, b io.Reader, impl excel.Utils) {
	fa, ha, err := impl.Read(a)
	So(err, ShouldBeNil)
	fb, hb, err := impl.Read(b)
	So(err, ShouldBeNil)
	So(fa, ShouldResemble, fb)
	So(ha, ShouldResemble, hb)
}

func TestCsvImpl(t *testing.T) {
	PatchConvey("case csvImpl", t, func() {
		input, ouput := "test.csv", "ans.csv"
		impl := excel.NewCsvImpl()
		f, err := os.Open(input)
		if err != nil {
			So(err, ShouldBeNil)
		}
		dataset, header, err := impl.Read(f)
		So(f.Close(), ShouldBeNil)
		So(err, ShouldBeNil)
		So(header, ShouldNotBeEmpty)
		So(dataset, ShouldNotBeEmpty)
		fmt.Println(dataset, header)
		fw, err := os.Create(ouput)
		if err != nil {
			So(err, ShouldBeNil)
		}
		So(impl.Write(fw, dataset, header), ShouldBeNil)
		fa, err := os.Open(input)
		So(err, ShouldBeNil)
		fb, err := os.Open(ouput)
		So(err, ShouldBeNil)
		compare2File(fa, fb, impl)
	})

	PatchConvey("case excelImpl", t, func() {
		input, ouput := "test.xlsx", "ans.xlsx"
		impl := excel.NewExcelImpl("")
		f, err := os.Open(input)
		if err != nil {
			So(err, ShouldBeNil)
		}
		dataset, header, err := impl.Read(f)
		So(f.Close(), ShouldBeNil)
		So(err, ShouldBeNil)
		So(header, ShouldNotBeEmpty)
		So(dataset, ShouldNotBeEmpty)
		fmt.Println(dataset, header)
		fw, err := os.Create(ouput)
		if err != nil {
			So(err, ShouldBeNil)
		}
		So(impl.Write(fw, dataset, header), ShouldBeNil)
		fa, err := os.Open(input)
		So(err, ShouldBeNil)
		fb, err := os.Open(ouput)
		So(err, ShouldBeNil)
		compare2File(fa, fb, impl)
	})
}
