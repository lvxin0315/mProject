package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/mProject/dictionaries"
	"github.com/lvxin0315/mProject/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Demo(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.html", nil)
}

func GetDemoData(c *gin.Context) {
	//cfg, err := goconfig.LoadConfigFile("etc/gorm.ini")
	//if err != nil {
	//	logrus.Error("goconfig.LoadConfigFile is error:", err)
	//	return
	//}
	//dbName, err := cfg.GetValue("gorm", "dbName")
	//if err != nil {
	//	logrus.Error("cfg.GetValue is error:", err)
	//	return
	//}
	//infoList, err := service.GetFiledByTableName(dbName, "h_litestore_goods")
	//if err != nil {
	//	c.HTML(http.StatusOK, "demo.html", nil)
	//	return
	//}
	//填充表头
	tableData := new(TableData)
	//记录顺序
	var fieldOrder []string
	for _, item := range dictionaries.YingShouKuanMingXi {
		tableData.Ths = append(tableData.Ths, item.Title)
		fieldOrder = append(fieldOrder, item.Field)
	}
	//数据填充
	tableData.Trs = [][]string{}

	dataList, err := service.GetAllDataMap()
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info(dataList)
	//整理成二维数组格式，按照字典顺序
	for _, item := range dataList {
		var tr []string
		for _, field := range fieldOrder {
			tr = append(tr, item[field])
		}
		tableData.Trs = append(tableData.Trs, tr)
	}
	c.JSON(http.StatusOK, tableData)
}
