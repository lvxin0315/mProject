package service

import (
	"github.com/lvxin0315/mProject/db_conn"
	"github.com/sirupsen/logrus"
)

type ColumnsInfo struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	DataType      string `gorm:"column:DATA_TYPE"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
}

func GetFiledByTableName(dbName, tableName string) ([]*ColumnsInfo, error) {
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var infoList []*ColumnsInfo
	sql := `Select COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT
from INFORMATION_SCHEMA.COLUMNS  
Where table_schema = ?
AND table_name = ? `
	db.Raw(sql, dbName, tableName).Scan(&infoList)
	return infoList, nil
}
