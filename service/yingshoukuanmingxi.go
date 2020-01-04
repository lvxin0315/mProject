package service

import (
	"encoding/json"
	"github.com/Chain-Zhang/pinyin"
	"github.com/lvxin0315/mProject/db_conn"
	"github.com/lvxin0315/mProject/model"
	"github.com/sirupsen/logrus"
)

func YinShouKuan(jsonContent map[string]interface{}) error {
	ysk := jsonContent["应收款"].([][]string)

	//字段与索引
	ik := make(map[int]string)
	for i, z := range ysk[0] {
		k, err := pinyin.New(z).Split("_").Mode(pinyin.WithoutTone).Convert()
		if err != nil {
			return err
		}
		ik[i] = k
	}

	var YinShouKuanData []map[string]string

	for _, yArray := range ysk[1:] {
		data := make(map[string]string)
		for i, v := range yArray {
			data[ik[i]] = v
		}
		YinShouKuanData = append(YinShouKuanData, data)
	}

	logrus.Info(YinShouKuanData)

	//通过json转成struct
	YinShouKuanByte, err := json.Marshal(YinShouKuanData)
	if err != nil {
		logrus.Error(err)
		return err
	}

	var YinShouKuanList []*model.YingShouKuanMingXi

	err = json.Unmarshal(YinShouKuanByte, &YinShouKuanList)
	if err != nil {
		logrus.Error(err)
		return err
	}

	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		return err
	}

	for _, yinShouKuan := range YinShouKuanList {
		db.Create(yinShouKuan)
		if yinShouKuan.ID <= 0 {
			logrus.Error(yinShouKuan.HeTongBianHao, " is Error")
		}
	}

	return nil
}
