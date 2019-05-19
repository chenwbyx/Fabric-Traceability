package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"github.com/hyperledger/fabric/protos/peer"
)

type EducationChaincode struct {

}

func (t *EducationChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func (t *EducationChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addCom"{
		return t.addCom(stub, args)
	}else if fun == "queryComByCertNoAndName" {
		return t.queryComByCertNoAndName(stub, args)
	}else if fun == "queryComInfoByEntityID" {
		return t.queryComInfoByEntityID(stub, args)
	}else if fun == "updateCom" {
		return t.updateCom(stub, args)
	}else if fun == "delCom"{
		return t.delCom(stub, args)
	}

	return shim.Error("指定的函数名称错误")

}

func main(){
	err := shim.Start(new(EducationChaincode))
	if err != nil{
		fmt.Printf("启动EducationChaincode时发生错误: %s", err)
	}
}
