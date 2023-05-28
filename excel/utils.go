package excel

import (
	"io"

	"github.com/eeianux/cutils/excel/excel_impl"

	"github.com/eeianux/cutils/excel/core"
	"github.com/eeianux/cutils/excel/csv_impl"
)

type Utils interface {
	Read(io.Reader) (core.DataSet, []string, error)
	Write(closer io.WriteCloser, data core.DataSet, header []string) error
}

func NewCsvImpl() Utils {
	return &csv_impl.CsvImpl{}
}

func NewExcelImpl(sheet string) Utils {
	if sheet == "" {
		sheet = "sheet1"
	}
	return &excel_impl.ExcelImpl{
		Sheet: sheet,
	}
}
