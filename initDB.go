package main

import (
	"github.com/lvxin0315/mProject/db_conn"
	"github.com/lvxin0315/mProject/model"
)

func main() {
	db, err := db_conn.GetGormDB()
	if err != nil {
		panic(err)
	}
	db.DropTableIfExists(
		//model.YingShouKuanMingXi{},
		model.ChanZhiDaCheng{},
	)
	db.CreateTable(
		//model.YingShouKuanMingXi{},
		model.ChanZhiDaCheng{},
	)
}
