package BEA

import (
	"encoding/hex"
	"fmt"

	"github.com/zhulingbiezhi/sdkOld/TLV"
)

type RefundTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewRefund(trans *TransactionData, config *Config) (*RefundTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
	return &RefundTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0200",
		processingCode: "200000",
	}, nil
}

func (refund *RefundTransaction) Valid() error {
	if err := refund.baseValid(); err != nil {
		return err
	}
	return validMatch(refund.transData.Pan,
		refund.transData.Amount,
		refund.transData.TransId,
		refund.transData.CardExpireDate,
		refund.transData.Track2,
	)
}

func (refund *RefundTransaction) SetFields() {
	refund.baseFieldSet()
	refund.set(0, refund.messageTypeId)
	refund.set(3, refund.processingCode)

	var de22 string

	switch refund.transData.PosEntryMode {
	case INSERT:
		de22 = "05"
	case SWIPE:
		de22 = "02"
	case MANUAL:
		de22 = "01"
	case FALLBACK:
		de22 = "80"
	case WAVE:
		de22 = "70"
	}

	if len(refund.transData.Pin) > 0 {
		de22 += "1"
	} else {
		de22 += "2"
	}

	refund.set(22, de22)
	refund.set(24, param[refund.transData.TransType].nii)
	refund.set(25, param[refund.transData.TransType].posCondictionCode)

	switch refund.transData.PosEntryMode {
	case INSERT:
		fallthrough
	case WAVE:
		var iccData = make(map[string]string)

		for _, tag := range DE55TagList {
			fmt.Printf("tag:%v\n", tag)
			iccData[tag] = refund.transData.IccRelatedData[tag]
		}

		de55 := TLV.BuildConstructTLVMsg(iccData)
		refund.set(55, hex.EncodeToString(de55))
	default:
	}

	refund.set(62, refund.transData.TransId)

}

func (refund *RefundTransaction) Fields() []uint8 {
	return refund.entryMap[refund.transData.PosEntryMode]
}

func (refund *RefundTransaction) Name() string {
	return string(refund.transData.TransType)
}
