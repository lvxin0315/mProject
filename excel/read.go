package excel

import (
	"github.com/tealeg/xlsx"
)

func ReadExcelFile(excelFileName string) (map[string]interface{}, error) {
	// 打开文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return nil, err
	}

	// 内容
	var jsonContent = make(map[string]interface{})
	// 遍历sheet页读取
	for _, sheet := range xlFile.Sheets {
		var sheetContent [][]string
		//遍历行读取
		for _, row := range sheet.Rows {
			var rowContent []string
			// 遍历每行的列读取
			for _, cell := range row.Cells {
				rowContent = append(rowContent, cell.String())
			}
			sheetContent = append(sheetContent, rowContent)
		}
		jsonContent[sheet.Name] = sheetContent
	}
	return jsonContent, nil
}
