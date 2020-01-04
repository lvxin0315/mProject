package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/mProject/controller"
	"github.com/sirupsen/logrus"
)

func main() {
	creeperApiEngine := gin.New()
	logrus.Info("0.0.0.0:8083")
	creeperApiEngine.Static("/static", "./static")

	creeperApiEngine.LoadHTMLGlob("./tpl/*")

	creeperApiEngine.Any("/demo", controller.Demo)
	creeperApiEngine.GET("/getDemoData", controller.GetDemoData)

	creeperApiEngine.Any("/upload", controller.UploadPage)
	creeperApiEngine.POST("/uploadyingshoukuanmingxi", controller.UploadYingShouKuanMingXi)
	creeperApiEngine.POST("/uploadchanzhidacheng", controller.UploadChanZhiDaCheng)

	_ = creeperApiEngine.Run(":8083")
}
