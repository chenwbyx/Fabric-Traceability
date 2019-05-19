/*
 *Author:chwenbo
 *2019-4-14
 */
package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"fmt"
	"bytes"
)

const DOC_TYPE = "comObj"
const DOC_COM_TYPE = "comObj"

func PutCom(stub shim.ChaincodeStubInterface, com Commodity) ([]byte, bool) {

	com.ObjectType = DOC_COM_TYPE

	b, err := json.Marshal(com)
	if err != nil {
		return nil, false
	}

	err = stub.PutState(com.Primarykey, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

func GetComInfo(stub shim.ChaincodeStubInterface, primarykey string) (Commodity, bool)  {
	var com Commodity

	b, err := stub.GetState(primarykey)
	if err != nil {
		return com, false
	}

	if b == nil {
		return com, false
	}

	err = json.Unmarshal(b, &com)
	if err != nil {
		return com, false
	}

	return com, true
}

func getComByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer  resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil

}

func (t *EducationChaincode) addCom(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var com Commodity
	err := json.Unmarshal([]byte(args[0]), &com)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	// 查重: 身份证号码必须唯一
	_, exist := GetComInfo(stub, com.Primarykey)
	if exist {
		return shim.Error("要添加的溯源ID已存在")
	}

	_, bl := PutCom(stub, com)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

// 暂时搁置
func (t *EducationChaincode) queryComByCertNoAndName(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}
	CertNo := args[0]
	name := args[1]

	// 拼装CouchDB所需要的查询字符串(是标准的一个JSON串)
	// queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"eduObj\", \"CertNo\":\"%s\"}}", CertNo)
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"%s\", \"CertNo\":\"%s\", \"Name\":\"%s\"}}", DOC_TYPE, CertNo, name)

	// 查询数据
	result, err := getComByQueryString(stub, queryString)
	if err != nil {
		return shim.Error("根据证书编号及姓名查询信息时发生错误")
	}
	if result == nil {
		return shim.Error("根据指定的证书编号及姓名没有查询到相关的信息")
	}
	return shim.Success(result)
}


func (t *EducationChaincode) queryComInfoByEntityID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据溯源ID查询信息失败")
	}

	if b == nil {
		return shim.Error("根据溯源ID没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var com Commodity
	err = json.Unmarshal(b, &com)
	if err != nil {
		return  shim.Error("反序列化溯源信息失败")
	}

	// 获取历史变更数据
	iterator, err := stub.GetHistoryForKey(com.Primarykey)
	if err != nil {
		return shim.Error("根据指定的身份证号码查询对应的历史变更数据失败")
	}
	defer iterator.Close()

	// 迭代处理
	var historys []HistoryItem
	var hisCom Commodity
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("获取edu的历史变更数据失败")
		}

		var historyItem HistoryItem
		historyItem.TxId = hisData.TxId
		json.Unmarshal(hisData.Value, &hisCom)

		if hisData.Value == nil {
			var empty Commodity
			historyItem.Commodity = empty
		}else {
			historyItem.Commodity = hisCom
		}

		historys = append(historys, historyItem)

	}

	com.Historys = historys

	result, err := json.Marshal(com)
	if err != nil {
		return shim.Error("序列化溯源信息时发生错误")
	}
	return shim.Success(result)
}

func (t *EducationChaincode) updateCom(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var info Commodity
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return  shim.Error("反序列化com信息失败")
	}

	result, bl := GetComInfo(stub, info.Primarykey)
	if !bl{
		return shim.Error("根据溯源编号查询信息时发生错误")
	}

	result.Type = info.Type
	result.Name = info.Name
	result.Des = info.Des
	result.Specification = info.Specification
	result.Source = info.Source
	result.Machining = info.Machining
	result.Photo = info.Photo
	result.Remarks = info.Remarks
	result.Principal = info.Principal
	result.PhoneNumber = info.PhoneNumber

	result.ShelfLife = info.ShelfLife
	result.StorageMethod = info.StorageMethod
	result.Brand = info.Brand
	result.Vendor = info.Vendor
	result.PlaceOfProduction = info.PlaceOfProduction
	result.ExecutiveStandard = info.ExecutiveStandard

	result.Time = info.Time


	_, bl = PutCom(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}

func (t *EducationChaincode) delCom(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	/*var edu Education
	result, bl := GetEduInfo(stub, info.EntityID)
	err := json.Unmarshal(result, &edu)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}*/

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("删除信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息删除成功"))
}

