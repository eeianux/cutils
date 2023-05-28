package test

import (
	"fmt"
	. "github.com/bytedance/mockey"
	"github.com/eeianux/cutils/excel_utils"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"os"
	"testing"
)

func compare2File(a, b io.Reader, impl excel_utils.Utils) {
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
		impl := excel_utils.NewCsvImpl()
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
		impl := excel_utils.NewExcelImpl("")
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
