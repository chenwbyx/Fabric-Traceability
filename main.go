package main

import (
	"os"
	"fmt"
	"time"
	"github.com/chenwbyx/Fabric-Traceability/sdkInit"
	"github.com/chenwbyx/Fabric-Traceability/service"
	"encoding/json"
	"github.com/chenwbyx/Fabric-Traceability/web/controller"
	"github.com/chenwbyx/Fabric-Traceability/web"
)

const (
	configFile = "config.yaml"
	initialized = false
	EduCC = "educc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID: "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chenwbyx/Fabric-Traceability/fixtures/artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID: EduCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "github.com/chenwbyx/Fabric-Traceability/chaincode/",
		UserName:"User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID:EduCC,
		Client:channelClient,
	}
	/*
	edu := service.Education{
		Name: "张三",
		Gender: "男",
		Nation: "汉",
		EntityID: "101",
		Place: "北京",
		BirthDay: "1991年01月01日",
		EnrollDate: "2009年9月",
		GraduationDate: "2013年7月",
		SchoolName: "中国政法大学",
		Major: "社会学",
		QuaType: "普通",
		Length: "四年",
		Mode: "普通全日制",
		Level: "本科",
		Graduation: "毕业",
		CertNo: "111",
		Photo: "/static/photo/11.png",
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
	*/
	com := service.Commodity{
		Type:"采摘",
		Primarykey:"001",
		Name:"普洱茶",
		Des:"从地里采摘",
		Specification:"500g",  //规格
		Source:"普洱",
		Machining:"人工采摘",    //加工
		Remarks:"无",    //备注信息
		Principal:"张三",  //负责人
		PhoneNumber:"123456789",
		Photo:"/static/photo/11.png",   // 照片
		ShelfLife:"一年",
		StorageMethod:"避光，常温",
		Brand:"普洱",
		Vendor:"云南某某厂",
		PlaceOfProduction:"云南",
		ExecutiveStandard:"GB/T 11766",
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}

	coms := []service.Commodity{
		service.Commodity{Type:"采摘", Primarykey:"002", Name:"铁观音", Des:"从地里采摘", Specification:"500g", Source:"普洱", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/11.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"采摘", Primarykey:"003", Name:"大红袍", Des:"从地里采摘", Specification:"500g", Source:"普洱", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/11.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"采摘", Primarykey:"004", Name:"小青柑", Des:"从地里采摘", Specification:"500g", Source:"普洱", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/11.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
	}

	for _,v := range coms {
		msg, err := serviceSetup.SaveCom(v)
		if err != nil {
			fmt.Println(err.Error())
		}else {
			fmt.Println("信息发布成功, 交易编号为: " + msg)
		}
	}
	/*
	msg, err := serviceSetup.SaveEdu(edu)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}
	*/
	msg, err := serviceSetup.SaveCom(com)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}
	/*
	result, err := serviceSetup.FindEduInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		//fmt.Println(edu)
	}
	*/
	result, err := serviceSetup.FindComInfoByEntityID("001")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
	}
	/*
	info := service.Education{
		Name: "张三",
		Gender: "男",
		Nation: "汉",
		EntityID: "101",
		Place: "北京",
		BirthDay: "1991年01月01日",
		EnrollDate: "2013年9月",
		GraduationDate: "2015年7月",
		SchoolName: "中国政法大学",
		Major: "社会学",
		QuaType: "普通",
		Length: "两年",
		Mode: "普通全日制",
		Level: "研究生",
		Graduation: "毕业",
		CertNo: "333",
		Photo: "/static/photo/11.png",
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
	msg, err = serviceSetup.ModifyEdu(info)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息操作成功, 交易编号为: " + msg)
	}

	result, err = serviceSetup.FindEduInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		//fmt.Println(edu)
	}

	result, err = serviceSetup.FindEduByCertNoAndName("333","张三")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		//fmt.Println(edu)
	}
	*/
	//===========================================//

	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)

}
