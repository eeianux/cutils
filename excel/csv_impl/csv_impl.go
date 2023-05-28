package csv_impl

import (
	"encoding/csv"
	"io"

	"github.com/eeianux/cutils/excel/core"
)

type CsvImpl struct{}

func (c CsvImpl) Read(reader io.Reader) (core.DataSet, []string, error) {
	csvReader := csv.NewReader(reader)
	content, err := csvReader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	data, header := core.ToDataSet(content)
	return data, header, nil
}

func (c CsvImpl) Write(closer io.WriteCloser, data core.DataSet, header []string) error {
	writer := csv.NewWriter(closer)
	content, _ := core.FromDataSet(data, header)
	return writer.WriteAll(content)
}
