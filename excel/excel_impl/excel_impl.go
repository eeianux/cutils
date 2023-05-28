package excel_impl

import (
	"io"

	"github.com/eeianux/cutils/excel/core"
	"github.com/eeianux/cutils/excel/utils"
	"github.com/xuri/excelize/v2"
)

type ExcelImpl struct {
	Sheet string
}

func (e ExcelImpl) Read(reader io.Reader) (core.DataSet, []string, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, nil, err
	}
	content, err := f.GetRows(e.Sheet)
	if err != nil {
		return nil, nil, err
	}
	data, header := core.ToDataSet(content)
	return data, header, nil
}

func (e ExcelImpl) Write(closer io.WriteCloser, data core.DataSet, header []string) error {
	f := excelize.NewFile()
	writer, err := f.NewStreamWriter(e.Sheet)
	if err != nil {
		return err
	}
	content, _ := core.FromDataSet(data, header)
	for idx, item := range content {
		rows, err := utils.ToInterfaceSlice(item)
		if err != nil {
			return err
		}
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			return err
		}
		if err := writer.SetRow(cell, rows); err != nil {
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return f.Write(closer)
}
