package BEA

import (
	"fmt"
)

/*
func DoRequest(transData *TransactionData, config *Config) (*TransactionData, error) {
	err := getSupportEntryMode(transData.PosEntryMode)
	if err != nil {
		return nil, fmt.Errorf("GetSupportEntryMode error: %s", err.Error())
	}

	fields, err := getFields(transData.TransType, transData.PosEntryMode)
	if err != nil {
		return nil, fmt.Errorf("GetFields error: %s", err.Error())
	}
	replyData, err := communicateWithHost(transData, config, fields)
	if err != nil {
		switch err {
		case RECV_ERR:
			transData.ResponseCode = BINDO_RECV_ERR
		case CONN_ERR:
			transData.ResponseCode = BINDO_CONN_ERR
		case SEND_ERR:
			transData.ResponseCode = BINDO_SEND_ERR
		default:
		}
		return nil, fmt.Errorf("communicateWithHost error: %s", err.Error())
	}
	return replyData, nil
}*/

func DoRequest(transData *TransactionData, config *Config) (*TransactionData, error) {
	//生成NewAuthorize结构体
	//Valid()
	//Fields()
	//删除多余的field
	//正式开始调用create message，进行交易
	if transData.TransType == "" {
		return nil, fmt.Errorf("the TransactionData.TransType is empty !")
	}
	txnInter, err := getTransactionInterface(transData, config)
	if err != nil {
		return nil, fmt.Errorf("getTransactionInterface %s error: %s", transData.TransType, err.Error())
	}
	if err := txnInter.Valid(); err != nil {
		return nil, err
	}
	txnInter.SetFields()
	fields := txnInter.Fields()
	fieldsMap := txnInter.GetFieldsMap()

	finalMap := make(map[uint8]string)
	for _, id := range fields {
		val := fieldsMap[id]
		if val == "" {
			return nil, fmt.Errorf("the TransactionInterface %s of field %d is empty", txnInter.Name(), id)
		} else {
			finalMap[id] = val
		}
	}
	replyData, err := communicateWithHost(transData, config, finalMap)
	if err != nil {
		switch err {
		case RECV_ERR:
			transData.ResponseCode = BINDO_RECV_ERR
		case CONN_ERR:
			transData.ResponseCode = BINDO_CONN_ERR
		case SEND_ERR:
			transData.ResponseCode = BINDO_SEND_ERR
		default:
		}
		return nil, fmt.Errorf("communicateWithHost error: %s", err.Error())
	}
	return replyData, nil
}

func getTransactionInterface(transData *TransactionData, config *Config) (BEAInterface, error) {
	switch transData.TransType {
	case KindPreAuthorize:
		return NewAuthorize(transData, config)
	default:
		return nil, fmt.Errorf("no match bea TransactionInterface")
	}
}
