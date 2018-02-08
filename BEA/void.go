package BEA

import (
	"encoding/hex"
	"fmt"

	"github.com/zhulingbiezhi/sdkOld/TLV"
)

type VoidTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewVoid(trans *TransactionData, config *Config) (*VoidTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
	trxn := &VoidTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0200",
		processingCode: "000000",
	}

	switch trans.OriginalTransType {
	case KindSale:
		trxn.processingCode = "020000"
	case KindRefund:
		trxn.processingCode = "220000"
	case KindPreAuthorize:
		trxn.processingCode = "000000"
	case KindPreAuthCompletion:
		trxn.processingCode = "000000"
	default:
		return nil, fmt.Errorf("unknow transaction type:%s", string(trans.OriginalTransType))
	}

	return trxn, nil
}

func (void *VoidTransaction) Valid() error {
	if err := void.baseValid(); err != nil {
		return err
	}
	return validMatch(void.transData.Pan,
		void.transData.Amount,
		void.transData.TransId,
		void.transData.CardExpireDate,
		void.transData.Track2,
	)
}

func (void *VoidTransaction) SetFields() {
	void.baseFieldSet()
	void.set(0, void.messageTypeId)
	void.set(3, void.processingCode)
	void.set(24, param[void.transData.TransType].nii)
	void.set(25, param[void.transData.TransType].posCondictionCode)

	var de22 string

	switch void.transData.PosEntryMode {
	case INSERT:
		de22 = "05"
	case SWIPE:
		de22 = "02"
	case MANUAL:
		de22 = "01"
	case WAVE:
		de22 = ""
	}

	if len(void.transData.Pin) > 0 {
		de22 += "1"
	} else {
		de22 += "2"
	}

	void.set(22, de22)
	void.set(24, param[void.transData.TransType].nii)
	void.set(25, param[void.transData.TransType].posCondictionCode)

	switch void.transData.PosEntryMode {
	case INSERT:
		fallthrough
	case WAVE:
		var iccData = make(map[string]string)

		for _, tag := range DE55TagList {
			fmt.Printf("tag:%v\n", tag)
			iccData[tag] = void.transData.IccRelatedData[tag]
		}

		de55 := TLV.BuildConstructTLVMsg(iccData)
		void.set(55, hex.EncodeToString(de55))
	default:
	}

	void.set(62, void.transData.TransId)
}

func (void *VoidTransaction) Fields() []uint8 {
	return void.entryMap[void.transData.PosEntryMode]
}

func (void *VoidTransaction) Name() string {
	return string(void.transData.TransType)
}
