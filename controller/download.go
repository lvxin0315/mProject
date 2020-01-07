package controller

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/mProject/excel"
	"github.com/lvxin0315/mProject/service"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func DownloadOne(c *gin.Context) {
	//获取数据
	resultDataList, err := service.GetAllChanZhiDaChengDataByQuYu()
	if err != nil {
		logrus.Error(err)
		return
	}
	//生成多个excel
	timeInt64 := time.Now().UnixNano()
	tmpPath := fmt.Sprintf("tmp/%d", timeInt64)
	err = os.Mkdir(tmpPath, 0777)
	if err != nil {
		logrus.Error(err)
		return
	}
	//添加区域报告内容并进行记录
	var zipFiles []string
	for quyu, dataList := range resultDataList {
		excelF := excelize.NewFile()
		sheetName := quyu + "报告"
		excelF.NewSheet(sheetName)
		//表头
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("L%d", 7), "项目名称")
		_ = excelF.SetColWidth(sheetName, "L", "L", 30)
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("M%d", 7), "项目经理")
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("N%d", 7), "里程碑")
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("O%d", 7), "月初计划产值（万元")
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("P%d", 7), "月初计划软件产值(万元)")
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("Q%d", 7), "项目进展")
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("R%d", 7), "达成情况")
		_ = excelF.SetColWidth(sheetName, "M", "R", 20)
		//设置表头背景颜色
		btStyle1, _ := excelF.NewStyle(`{"fill":{"type":"pattern","color":["#C0D9EF"],"pattern":1}}`)
		_ = excelF.SetCellStyle(sheetName, "L7", "P7", btStyle1)
		btStyle2, _ := excelF.NewStyle(`{"fill":{"type":"pattern","color":["#FFD4B7"],"pattern":1}}`)
		_ = excelF.SetCellStyle(sheetName, "Q7", "R7", btStyle2)
		//数据
		for row, rData := range dataList {
			_ = excelF.SetCellValue(sheetName, fmt.Sprintf("L%d", row+8), rData.XiangMuMingCheng)
			_ = excelF.SetCellValue(sheetName, fmt.Sprintf("M%d", row+8), rData.XiangMuJingLi)
			_ = excelF.SetCellValue(sheetName, fmt.Sprintf("N%d", row+8), rData.LiChengBeiMingCheng)
			_ = excelF.SetCellValue(sheetName, fmt.Sprintf("O%d", row+8), excel.StringToFloat(rData.LiChengBeiDaChengChanZhi))
			_ = excelF.SetCellValue(sheetName, fmt.Sprintf("P%d", row+8), excel.StringToFloat(rData.LiChengBeiDaChengRuanJianChanZhi))
			_ = excelF.SetCellValue(sheetName, fmt.Sprintf("Q%d", row+8), "")
			_ = excelF.SetCellValue(sheetName, fmt.Sprintf("R%d", row+8), "")
		}

		//绘制统计图
		ec := excel.InitFormatChart()
		ec.Title.Name = quyu + "图表"
		ec.Type = "col"
		ec.Legend.Position = "bottom"
		//统计图数据
		ecCategories := fmt.Sprintf("%s!$M$%d:$M%d", sheetName, 8, len(dataList)+8-1)
		s1 := excel.FormatChartSeries{
			Name:       fmt.Sprintf("%s!$O7", sheetName),
			Categories: ecCategories,
			Values:     fmt.Sprintf("%s!$O$%d:$O$%d", sheetName, 8, len(dataList)+8-1),
		}

		s2 := excel.FormatChartSeries{
			Name:       fmt.Sprintf("%s!$P7", sheetName),
			Categories: ecCategories,
			Values:     fmt.Sprintf("%s!$P$%d:$P$%d", sheetName, 8, len(dataList)+8-1),
		}

		ec.Series = append(ec.Series, s1, s2)

		//填充统计图
		chartJson := excel.ChartStructToString(ec)
		if chartJson != "" {
			//logrus.Info("sheetName:", sheetName, "ec:",excel.ChartStructToString(ec))
			_ = excelF.AddChart(sheetName, "C10", excel.ChartStructToString(ec))
		}

		logrus.Info(excelF.GetPicture(sheetName, "C10"))

		saveFileName := fmt.Sprintf("%s/%s.xlsx", tmpPath, quyu)
		if err := excelF.SaveAs(saveFileName); err != nil {
			logrus.Error(saveFileName, " is error: ", err)
			return
		}
		zipFiles = append(zipFiles, saveFileName)
	}
	logrus.Info(zipFiles)
	//zip压缩
	downFile := fmt.Sprintf("%s/%s.zip", tmpPath, "new")
	err = excel.Zip(downFile, zipFiles)
	if err != nil {
		logrus.Error("excel zip error:", err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "new.zip"))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(downFile)
}
