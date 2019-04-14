package service

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"fmt"
	"time"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type Education struct {
	ObjectType	string	`json:"docType"`
	Name	string	`json:"Name"`
	Gender	string	`json:"Gender"`
	Nation	string	`json:"Nation"`
	EntityID	string	`json:"EntityID"`
	Place	string	`json:"Place"`
	BirthDay	string	`json:"BirthDay"`
	EnrollDate	string	`json:"EnrollDate"`
	GraduationDate	string	`json:"GraduationDate"`
	SchoolName	string	`json:"SchoolName"`
	Major	string	`json:"Major"`
	QuaType	string	`json:"QuaType"`
	Length	string	`json:"Length"`
	Mode	string	`json:"Mode"`
	Level	string	`json:"Level"`
	Graduation	string	`json:"Graduation"`
	CertNo	string	`json:"CertNo"`

	Photo	string	`json:"Photo"`

	Historys	[]HistoryItem
	Time    string  `json:"Time"`
}

type Commodity struct {
	ObjectType	   string	`json:"docType"`
	Primarykey     string   `json:"primarykey"`  //主键，唯一Id
	Name	       string	`json:"Name"`
	Type           string   `json:"type"`
	Des            string   `json:"des"`
	Specification  string   `json:"specification"`  //规格
	Source         string   `json:"source"`
	Machining      string   `json:"machining"`    //加工
	Remarks        string   `json:"remarks"`    //备注信息
	Principal      string   `json:"principal"`  //负责人
	PhoneNumber    string   `json:"phoneNumber"`
	Photo	string	`json:"Photo"`	            // 照片
	Historys	[]HistoryItem	// 当前edu的历史记录
	Time    string  `json:"Time"`   //时间
}

type HistoryItem struct {
	TxId	string
	Education	Education
	Commodity   Commodity
}

type ServiceSetup struct {
	ChaincodeID	string
	Client	*channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
