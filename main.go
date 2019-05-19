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
	ComCC = "comcc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID: "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chenwbyx/Fabric-Traceability/fixtures/artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID: ComCC,
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
		ChaincodeID:ComCC,
		Client:channelClient,
	}

	coms := []service.Commodity{
		service.Commodity{Type:"采摘", Primarykey:"001", Name:"普洱茶", Des:"从地里采摘", Specification:"500g", Source:"普洱", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/11.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
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

	result, err := serviceSetup.FindComInfoByEntityID("001")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var com service.Commodity
		json.Unmarshal(result, &com)
	}

	coms = []service.Commodity{
		service.Commodity{Type:"加工", Primarykey:"001", Name:"普洱茶", Des:"从地里采摘", Specification:"500g", Source:"普洱", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/11.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"加工", Primarykey:"002", Name:"铁观音", Des:"从地里采摘", Specification:"500g", Source:"普洱", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/11.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"加工", Primarykey:"003", Name:"大红袍", Des:"从地里采摘", Specification:"500g", Source:"普洱", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/11.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
		service.Commodity{Type:"加工", Primarykey:"004", Name:"小青柑", Des:"从地里采摘", Specification:"500g", Source:"普洱", Machining:"人工采摘",Remarks:"无",Principal:"张三",PhoneNumber:"123456789", Photo:"/static/photo/11.png", ShelfLife:"一年", StorageMethod:"避光，常温", Brand:"普洱", Vendor:"云南某某厂", PlaceOfProduction:"云南", ExecutiveStandard:"GB/T 11766", Time: time.Now().Format("2006-01-02 15:04:05"),},
	}

	for _,v := range coms {
		msg, err := serviceSetup.ModifyCom(v)
		if err != nil {
			fmt.Println(err.Error())
		}else {
			fmt.Println("信息操作成功, 交易编号为: " + msg)
		}
	}

	result, err = serviceSetup.FindComInfoByEntityID("001")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var com service.Commodity
		json.Unmarshal(result, &com)
		fmt.Println(com)
	}

	//===========================================//

	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)

}
