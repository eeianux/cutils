package core

type DataSet []map[string]string

func ToDataSet(content [][]string) (DataSet, []string) {
	var dataSet DataSet
	var header []string
	for idx, item := range content {
		if idx == 0 {
			for _, jdx := range item {
				header = append(header, jdx)
			}
			continue
		}
		d := make(map[string]string)
		for jdx, jtem := range item {
			d[header[jdx]] = jtem
		}
		dataSet = append(dataSet, d)
	}
	return dataSet, header
}

func FromDataSet(set DataSet, header []string) ([][]string, []string) {
	var content [][]string
	for idx, item := range set {
		if idx == 0 {
			if len(header) == 0 {
				for _, jtem := range item {
					header = append(header, jtem)
				}
			}
			content = append(content, header)
		}
		var data []string
		for _, field := range header {
			data = append(data, item[field])
		}
		content = append(content, data)
	}
	return content, header
}
