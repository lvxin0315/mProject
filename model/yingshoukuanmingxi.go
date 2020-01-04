package model

import "github.com/jinzhu/gorm"

//应收款明细
type YingShouKuanMingXi struct {
	gorm.Model
	QuYu                     string `json:"qu_yu"`
	HangYeXian               string `json:"hang_ye_xian"`
	QianDingNianFen          string `json:"qian_ding_nian_fen"`
	HeTongBianHao            string `json:"he_tong_bian_hao"`
	XiangMuMingCheng         string `json:"xiang_mu_ming_cheng"`
	HeTongJiaFang            string `json:"he_tong_jia_fang"`
	HeTongZongE              string `json:"he_tong_zong_e"`
	RuanJianHeTongE          string `json:"ruan_jian_he_tong_e"`
	ShiJiHuiKuan             string `json:"shi_ji_hui_kuan"`
	HuiKuanWanChengBi        string `json:"hui_kuan_wan_cheng_bi"`
	YingShouKuanHeJi         string `json:"ying_shou_kuan_he_ji"`
	RuanJianYingShouKuanHeJi string `json:"ruan_jian_ying_shou_kuan_he_ji"`
	QianKuanDengJi           string `json:"qian_kuan_deng_ji"`
	XiaoShouRenYuan          string `json:"xiao_shou_ren_yuan"`
}

func (m *YingShouKuanMingXi) TableName() string {
	return "YingShouKuanMingXi"
}
