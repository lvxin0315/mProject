package controller

import (
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/mProject/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Demo(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.html", nil)
}

func GetDemoData(c *gin.Context) {
	cfg, err := goconfig.LoadConfigFile("etc/gorm.ini")
	if err != nil {
		logrus.Error("goconfig.LoadConfigFile is error:", err)
		return
	}
	dbName, err := cfg.GetValue("gorm", "dbName")
	if err != nil {
		logrus.Error("cfg.GetValue is error:", err)
		return
	}
	infoList, err := service.GetFiledByTableName(dbName, "h_litestore_goods")
	if err != nil {
		c.HTML(http.StatusOK, "demo.html", nil)
		return
	}
	//填充表头
	tableData := new(TableData)
	for _, item := range infoList {
		tableData.Ths = append(tableData.Ths, item.ColumnComment)
	}
	//数据填充
	tableData.Trs = [][]string{}

	c.JSON(http.StatusOK, tableData)
}
