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

	_ = creeperApiEngine.Run(":8083")
}
