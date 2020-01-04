package main

import (
	"fmt"
	"github.com/Chain-Zhang/pinyin"
	"github.com/lvxin0315/mProject/excel"
)

func main() {
	data, _ := excel.ReadExcelFile("tmp/1.xlsx")
	ysk := data["1月"].([][]string)

	for _, z := range ysk[0] {
		d, _ := pinyin.New(z).Split("").Mode(pinyin.InitialsInCapitals).Convert()
		k, err := pinyin.New(z).Split("_").Mode(pinyin.WithoutTone).Convert()
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(fmt.Sprintf("{Title: \"%s\", Field: \"%s\",},\r\n", z, k))
		fmt.Println(fmt.Sprintf("%s  string `json:\"%s\"`\r\n", d, k))
	}
}
