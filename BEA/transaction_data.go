package BEA

import (
	"encoding/json"
	"fmt"
)

type BatchTotals struct {
	CapturedSalesCount    int
	CapturedSalesAmount   int64
	CapturedRefundCount   int
	CapturedRefundAmount  int64
	DebitSalesCount       int
	DebitSalesAmount      int64
	DebitRefundCount      int
	DebitRefundAmount     int64
	AuthorizeSalesCount   int
	AuthorizeSalesAmount  int64
	AuthorizeRefundCount  int
	AuthorizeRefundAmount int64
}

type TransactionData struct {
	TransId           string            `json:"trans_id"`             //流水号--------11
	AcquireTransID    string            `json:"acq_trans_id"`         //收单行交易号---37
	TransDate         string            `json:"trans_date"`           //交易日期------13
	TransTime         string            `json:"trans_time"`           //交易时间------12
	Amount            string            `json:"amount"`               //授权金额------04
	TipAmount         string            `json:"tip"`                  //消费金额
	Pin               string            `json:"pin"`                  //联机PINBLOCK--52
	Pan               string            `json:"pan"`                  //主账号--------02
	PanSeqNo          string            `json:"pan_seq_no,omitempty"` //卡片序列号-----
	CardExpireDate    string            `json:"card_exp_date"`        //有效期--------14
	Track1            string            `json:"track1,omitempty"`     //磁道一--------
	Track2            string            `json:"track2"`               //磁道二--------35
	PosEntryMode      EntryMode         `json:"pos_entry_mode"`       //刷卡方式------22
	IccRelatedData    map[string]string //IC卡相关数据--
	AuthCode          string            `json:"auth_code"`               //授权码-------38
	ResponseCode      BEACode           `json:"response_code"`           //响应码-------39
	Invoice           string            `json:"invoice,omitempty"`       //发票号-------62
	Batchtotals       BatchTotals       `json:"batch_total,omitempty"`   //settlement总金额----62
	BatchNumber       string            `json:"batch_number,omitempty"`  //settlement批次------60
	OriginalAmount    string            `json:"origin_amount,omitempty"` //原交易金额----
	OriginalTransType TransactionType   `json:"origin_amount,omitempty"` //原交易金额----
	TransType         TransactionType   `json:"trans_type,omitempty"`    //交易类型------
}

func (t TransactionData) FormJson() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("json marshal error: %s", err.Error())
	} else {
		return string(data)
	}
}

//后台配置参数
type Config struct {
	Host       string //后台地址
	TPDU       string //tpdu
	EDS        string //eds
	TerminalId string
	MerchantId string
	TMK        string
	TMKIndex   string
	TimeOut    int
}
