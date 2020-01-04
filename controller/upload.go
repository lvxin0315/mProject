package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/mProject/excel"
	"github.com/lvxin0315/mProject/service"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

func UploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func UploadYingShouKuanMingXi(c *gin.Context) {
	file, header, err := c.Request.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	//文件的名称
	filename := header.Filename
	logrus.Println(file, err, filename)
	//创建文件
	timeInt64 := time.Now().UnixNano()
	err = os.Mkdir(fmt.Sprintf("tmp/%d", timeInt64), 0777)
	if err != nil {
		logrus.Error(err)
		return
	}
	newFile := fmt.Sprintf("tmp/%d/%s", timeInt64, filename)
	out, err := os.Create(newFile)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		logrus.Error(err)
		return
	}
	//读取excel内容

	jsonContent, err := excel.ReadExcelFile(newFile)
	if err != nil {
		logrus.Error("ReadExcelFile:", err)
		return
	}

	err = service.YinShouKuan(jsonContent)
	if err != nil {
		logrus.Error("YinShouKuan:", err)
		return
	}

	c.String(http.StatusOK, "upload successful")
}

func UploadChanZhiDaCheng(c *gin.Context) {
	file, header, err := c.Request.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	//文件的名称
	filename := header.Filename
	logrus.Println(file, err, filename)
	//创建文件
	timeInt64 := time.Now().UnixNano()
	err = os.Mkdir(fmt.Sprintf("tmp/%d", timeInt64), 0777)
	if err != nil {
		logrus.Error(err)
		return
	}
	newFile := fmt.Sprintf("tmp/%d/%s", timeInt64, filename)
	out, err := os.Create(newFile)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		logrus.Error(err)
		return
	}
	//读取excel内容

	jsonContent, err := excel.ReadExcelFile(newFile)
	if err != nil {
		logrus.Error("ReadExcelFile:", err)
		return
	}

	err = service.ChanZhiDaCheng(jsonContent)
	if err != nil {
		logrus.Error("ChanZhiDaCheng:", err)
		return
	}

	c.String(http.StatusOK, "upload successful")
}
