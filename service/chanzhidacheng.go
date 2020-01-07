package service

import (
	"encoding/json"
	"github.com/Chain-Zhang/pinyin"
	"github.com/lvxin0315/mProject/db_conn"
	"github.com/lvxin0315/mProject/model"
	"github.com/sirupsen/logrus"
	"strings"
)

func ChanZhiDaCheng(jsonContent map[string]interface{}) error {
	//TODO
	ysk := jsonContent["1月"].([][]string)

	//字段与索引
	ik := make(map[int]string)
	for i, z := range ysk[0] {
		k, err := pinyin.New(z).Split("_").Mode(pinyin.WithoutTone).Convert()
		if err != nil {
			return err
		}
		ik[i] = strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					k, "/", ""),
				"(", ""),
			")", "")

		//fmt.Println(fmt.Sprintf("{Title: \"%s\", Field: \"%s\",},\r\n", z, k))
	}

	logrus.Info(ik)

	var ChanZhiDaChengData []map[string]string

	for _, yArray := range ysk[1:] {
		data := make(map[string]string)
		for i, v := range yArray {
			data[ik[i]] = v
		}
		ChanZhiDaChengData = append(ChanZhiDaChengData, data)
	}

	logrus.Info(ChanZhiDaChengData)

	//通过json转成struct
	ChanZhiDaChengByte, err := json.Marshal(ChanZhiDaChengData)
	if err != nil {
		logrus.Error(err)
		return err
	}

	var ChanZhiDaChengList []*model.ChanZhiDaCheng

	logrus.Println(string(ChanZhiDaChengByte))

	err = json.Unmarshal(ChanZhiDaChengByte, &ChanZhiDaChengList)
	if err != nil {
		logrus.Error(err)
		return err
	}

	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		return err
	}

	for _, ChanZhiDaCheng := range ChanZhiDaChengList {
		db.Create(ChanZhiDaCheng)
		if ChanZhiDaCheng.ID <= 0 {
			logrus.Error(ChanZhiDaCheng.HeTongBianHao, " is Error")
		}
	}

	return nil
}

func GetAllChanZhiDaChengDataMap() ([]map[string]string, error) {
	dataList, err := GetAllChanZhiDaChengData()
	if err != nil {
		return nil, err
	}
	//使用json转map
	dataByte, _ := json.Marshal(dataList)
	var dataMap []map[string]string
	_ = json.Unmarshal(dataByte, &dataMap)
	return dataMap, nil
}

func GetAllChanZhiDaChengData() ([]*model.ChanZhiDaCheng, error) {
	var dataList []*model.ChanZhiDaCheng
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	db.Model(model.ChanZhiDaCheng{}).Find(&dataList)
	return dataList, nil
}

func GetAllChanZhiDaChengDataByQuYu() (map[string][]*model.ChanZhiDaCheng, error) {
	dataList, err := GetAllChanZhiDaChengData()
	if err != nil {
		return nil, err
	}
	resultDataList := make(map[string][]*model.ChanZhiDaCheng)
	for _, item := range dataList {
		resultDataList[item.QuYu] = append(resultDataList[item.QuYu], item)
	}
	return resultDataList, nil
}
