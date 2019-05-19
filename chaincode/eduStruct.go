package main

/**
姓名：张小三，性别：男，

民族：汉，身份证号：1011010101010101

籍贯：XXX，出生日期：1991年01月01日，			照片，

入学日期：2009年9月，毕（结）业日期：2013年7月，

学校名称：中国政法大学，专业：民商法学，

学历类别：普通，学制：四年，

学习形式：普通全日制，层次：本科，

毕（结）业：毕业，证书编号：11111111111111

 */

 /**
 商品唯一ID(溯源编号)：

 事件类型：

 简介：

 商品名称：

 商品规格：

 商品来源：

 加工方式：

 照片：

 备注信息：

 负责人：

 联系方式：

 录入时间：
  */
type Commodity struct {
	ObjectType	   string	`json:"docType"`
	Type           string	`json:"type"`  //事件类型
	Primarykey     string   `json:"primarykey"`  //主键，唯一Id
	Name	       string	`json:"name"`
	Des            string   `json:"des"`  //描述
	Specification  string   `json:"specification"`  //规格
	Source         string   `json:"source"`  //商品来源
	Machining      string   `json:"machining"`    //加工
	Remarks        string   `json:"remarks"`    //备注信息
	Principal      string   `json:"principal"`  //负责人
	PhoneNumber    string   `json:"phoneNumber"`
	Photo	       string	`json:"Photo"`	            // 照片

	ShelfLife      string   `json:"shelfLife"`   //保质期
	StorageMethod  string   `json:"storageMethod"`  //储藏方式
	Brand          string   `json:"brand"`  //品牌
	Vendor         string   `json:"vendor"`  //生产厂商
	PlaceOfProduction   string   `json:"placeOfProduction"`  //生产地
	ExecutiveStandard   string   `json:"executiveStandard"`  //执行标准

 	Historys	   []HistoryItem	// 当前com的历史记录
	Time           string  `json:"Time"`   //时间
}


type HistoryItem struct {
	TxId	string
	Commodity   Commodity
}
