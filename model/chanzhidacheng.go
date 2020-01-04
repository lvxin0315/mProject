package model

import "github.com/jinzhu/gorm"

type ChanZhiDaCheng struct {
	gorm.Model

	YueFen string `json:"yue_fen"`

	HeTongBianHao string `json:"he_tong_bian_hao"`

	XiangMuMingCheng string `json:"xiang_mu_ming_cheng"`

	XingYeXian string `json:"xing_ye_xian"`

	XiangMuJingLi string `json:"xiang_mu_jing_li"`

	LiChengBeiMingCheng string `json:"li_cheng_bei_ming_cheng"`

	LiChengBeiJiHuaDaChengShiJian string `json:"li_cheng_bei_ji_hua_da_cheng_shi_jian"`

	LiChengBeiShiJiWanChengShiJian string `json:"li_cheng_bei_shi_ji_wan_cheng_shi_jian"`

	LiChengBeiDaChengChanZhi string `gorm:"column:li_cheng_bei_da_cheng_chan_zhi__wan_yuan_" json:"li_cheng_bei_da_cheng_chan_zhi__wan_yuan_"`

	LiChengBeiDaChengRuanJianChanZhi string `gorm:"column:li_cheng_bei_da_cheng_ruan_jian_chan_zhi__wan_yuan_" json:"li_cheng_bei_da_cheng_ruan_jian_chan_zhi__wan_yuan_"`

	DaChengZhuangTai string `json:"da_cheng_zhuang_tai"`

	ShiJiHuiKuanJinE string `json:"shi_ji_hui_kuan_jin_e"`

	ShiJiHuiKuanShiJian string `json:"shi_ji_hui_kuan_shi_jian"`

	QianShuShiJian string `json:"qian_shu_shi_jian"`

	DangNianWangNian string `gorm:"column:dang_nian__wang_nian" json:"dang_nian__wang_nian"`

	XiaoShou string `json:"xiao_shou"`

	XiaoShouFuZeRen string `json:"xiao_shou_fu_ze_ren"`

	QuYu string `json:"qu_yu"`
}

func (m *ChanZhiDaCheng) TableName() string {
	return "ChanZhiDaCheng"
}
