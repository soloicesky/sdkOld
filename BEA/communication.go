package BEA

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/zhulingbiezhi/sdkOld/ISO8583"
	"github.com/zhulingbiezhi/sdkOld/TLV"
)

//和后台通信发送授权请求报文并接收授权响应报文
func sendReceiveData(reqMsg []byte, config *Config) ([]byte, error) {
	rspMsg := make([]byte, 0)
	conn, err := net.Dial("tcp", config.Host)
	if err != nil {
		return nil, CONN_ERR
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(time.Duration(config.TimeOut) * time.Second))
	count, err := conn.Write(reqMsg)
	if err != nil {
		return nil, SEND_ERR
	}

	totalLen := 0 //保存数据长度
	buf := make([]byte, 128)

	for {
		count, err = conn.Read(buf)
		if err != nil {
			return nil, RECV_ERR
		}

		rspMsg = append(rspMsg, buf[0:count]...)

		if len(rspMsg) >= 2 {
			totalLen = 2 + (int(rspMsg[0]) << 8) + int(rspMsg[1])
		}

		if totalLen > 0 && len(rspMsg) >= totalLen {
			break
		}
	}
	return rspMsg, nil
}

/**
	保存数据
	fieldId 域标识
	value  值
	storage 存储位置
**/
func saveData(fieldId int, value string, storage interface{}) error {
	transData, OK := storage.(*TransactionData)
	if !OK {
		return errors.New("interface is not a type of TransactionData")
	}

	fmt.Printf("id:%d value:%s\r\n", fieldId, value)

	switch fieldId {
	case 37:
		transData.AcquireTransID = value
	case 38:
		transData.AuthCode = value
	case 39:
		transData.ResponseCode = BEACode(value)
	case 55:
		de55, _ := hex.DecodeString(value)
		TLV.ParseConstructTLVMsg(de55, transData.IccRelatedData)
	}

	return nil
}

/**
	创建并发送一个授权报文
	transData 交易数据
	config 配置参数
	fields 域集合
**/
func communicateWithHost(transData *TransactionData, config *Config, fieldsMap map[uint8]string) (*TransactionData, error) {

	switch config.TerminalId {
	default:
		fallthrough
	case "63150001":
		msg, err := createIISO8583Message(fieldsMap, config)
		if err != nil {
			return transData, fmt.Errorf("CreateIISO8583Message error: %s", err.Error())
		}

		fmt.Printf("Final Msg:%s\r\n", hex.EncodeToString(msg))
		msg, err = sendReceiveData(msg, config)
		if err != nil {
			return nil, err
		}

		fmt.Printf("reponse ISO8583:%s\r\n", hex.EncodeToString(msg))
		err = ISO8583.DecodeISO8583Message(msg[2+5:], saveData, transData)
		if err != nil {
			transData.ResponseCode = BINDO_RECV_ERR
			return nil, fmt.Errorf("ISO8583::DecodeISO8583Message error: %s", err.Error())
		}
	case "63150002" //帐不平
		switch transData.TransType {
		case KindSettlment:
			transData.ResponseCode = "95"
		default:
			transData.ResponseCode = "00"
		}
		rrn := RandomStr(RandomStrTypeNumber, 12)
		transData.AcquireTransID = rrn
		authCode := RandomStr(RandomStrTypeNumber, 6)
		transData.AuthCode = authCode
	case "63150003" //帐平
		transData.ResponseCode = "00"
		rrn := RandomStr(RandomStrTypeNumber, 12)
		transData.AcquireTransID = rrn
		authCode := RandomStr(RandomStrTypeNumber, 6)
		transData.AuthCode = authCode
	}
	
	return transData, nil
}

const (
	RandomStrTypeNumber = 1
	RandomStrTypeAlpha  = 2
	RandomStrTypeMxied  = 3
)

var randomStrFull = "0123456789qwertyuiopasdfghjklzxcvbnm"
var r *rand.Rand = rand.New(rand.NewSource(time.Now().Unix()))

func RandomStr(t int, l int) string {

	min := 0
	max := len(randomStrFull)
	switch t {
	case RandomStrTypeNumber:
		max = 10
		break
	case RandomStrTypeAlpha:
		min = 10
		break
	case RandomStrTypeMxied:
		break
	}
	str := ""
	for i := 0; i < l; i++ {
		str += string(randomStrFull[min+r.Intn(max-min)])
	}
	return str
}
