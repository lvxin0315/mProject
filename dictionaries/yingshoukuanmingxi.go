package dictionaries

type Info struct {
	Title string
	Field string
}

var YingShouKuanMingXi = []*Info{
	{Title: "区域", Field: "qu_yu"},
	{Title: "行业线", Field: "xing_ye_xian"},
	{Title: "签订年份", Field: "qian_ding_nian_fen"},
	{Title: "合同编号", Field: "he_tong_bian_hao"},
	{Title: "项目名称", Field: "xiang_mu_ming_cheng"},
	{Title: "合同甲方", Field: "he_tong_jia_fang"},
	{Title: "合同总额", Field: "he_tong_zong_e"},
	{Title: "软件合同额", Field: "ruan_jian_he_tong_e"},
	{Title: "实际回款", Field: "shi_ji_hui_kuan"},
	{Title: "回款完成比", Field: "hui_kuan_wan_cheng_bi"},
	{Title: "应收款合计", Field: "ying_shou_kuan_he_ji"},
	{Title: "软件应收款合计", Field: "ruan_jian_ying_shou_kuan_he_ji"},
	{Title: "欠款等级", Field: "qian_kuan_deng_ji"},
	{Title: "销售人员", Field: "xiao_shou_ren_yuan"},
}
