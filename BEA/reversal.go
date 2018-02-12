package BEA

import (
	"encoding/hex"
	"fmt"

	"github.com/zhulingbiezhi/sdkOld/TLV"
)

type ReversalTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewReversal(trans *TransactionData, config *Config) (*ReversalTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}

	txn := &ReversalTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0400",
		processingCode: "000000",
	}

	switch trans.OriginalTransType {
	case KindSale:
		txn.processingCode = "000000"
	case KindVoid:
		txn.processingCode = "020000"
	case KindRefund:
		txn.processingCode = "220000"
	default:
		return nil, fmt.Errorf("unknow transaction type:%s", string(trans.OriginalTransType))
	}

	return txn, nil
}

func (reversal *ReversalTransaction) Valid() error {
	if err := reversal.baseValid(); err != nil {
		return err
	}
	return validMatch(reversal.transData.Pan,
		reversal.transData.Amount,
		reversal.transData.TransId,
		reversal.transData.CardExpireDate,
		// reversal.transData.Track2,
	)
}

func (reversal *ReversalTransaction) SetFields() {
	reversal.baseFieldSet()
	reversal.set(0, reversal.messageTypeId)
	reversal.set(3, reversal.processingCode)
	reversal.set(24, param[reversal.transData.TransType].nii)
	reversal.set(25, param[reversal.transData.TransType].posCondictionCode)
	var de22 string

	switch reversal.transData.PosEntryMode {
	case INSERT:
		de22 = "05"
	case SWIPE:
		de22 = "90"
	case MANUAL:
		de22 = "01"
	case FALLBACK:
		de22 = "80"
	case WAVE:
		de22 = "07"
	}

	if len(reversal.transData.Pin) > 0 {
		de22 += "1"
	} else {
		de22 += "2"
	}

	reversal.set(22, de22)
	reversal.set(24, param[reversal.transData.TransType].nii)
	reversal.set(25, param[reversal.transData.TransType].posCondictionCode)

	switch reversal.transData.PosEntryMode {
	case INSERT:
		fallthrough
	case WAVE:
		var iccData = make(map[string]string)

		for _, tag := range DE55TagList {
			// fmt.Printf("tag:%v\n", tag)
			iccData[tag] = reversal.transData.IccRelatedData[tag]
		}

		fmt.Printf("iccdata:%v\r\n", iccData)

		de55 := TLV.BuildConstructTLVMsg(iccData)
		reversal.set(55, hex.EncodeToString(de55))
	default:
	}

	reversal.set(62, reversal.transData.TransId)
}

func (reversal *ReversalTransaction) Fields() []uint8 {
	return reversal.entryMap[reversal.transData.PosEntryMode]
}

func (reversal *ReversalTransaction) Name() string {
	return string(reversal.transData.TransType)
}
