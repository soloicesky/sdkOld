package BEA

import (
	"encoding/hex"
	"fmt"

	"github.com/zhulingbiezhi/sdkOld/TLV"
)

type SaleTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewSale(trans *TransactionData, config *Config) (*SaleTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}

	return &SaleTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0200",
		processingCode: "000000",
	}, nil
}

func (sale *SaleTransaction) Valid() error {
	if err := sale.baseValid(); err != nil {
		return err
	}
	return validMatch(sale.transData.Pan,
		sale.transData.Amount,
		sale.transData.TransId,
		sale.transData.CardExpireDate,
		// sale.transData.Track2,
	)
}

func (sale *SaleTransaction) SetFields() {
	sale.baseFieldSet()
	sale.set(0, sale.messageTypeId)
	sale.set(3, sale.processingCode)

	var de22 string

	switch sale.transData.PosEntryMode {
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
	case MSD:
		de22 = "91"
	}

	if len(sale.transData.Pin) > 0 {
		de22 += "1"
	} else {
		de22 += "2"
	}

	sale.set(22, de22)
	sale.set(24, param[sale.transData.TransType].nii)
	sale.set(25, param[sale.transData.TransType].posCondictionCode)

	switch sale.transData.PosEntryMode {
	case INSERT:
		fallthrough
	case WAVE:
		var iccData = make(map[string]string)

		for _, tag := range DE55TagList {
			// fmt.Printf("tag:%v\n", tag)
			iccData[tag] = sale.transData.IccRelatedData[tag]
		}

		fmt.Printf("iccdata:%v\r\n", iccData)

		de55 := TLV.BuildConstructTLVMsg(iccData)
		sale.set(55, hex.EncodeToString(de55))
	default:
	}

	sale.set(54, fmt.Sprintf("%012s", sale.transData.TipAmount))
	sale.set(62, sale.transData.TransId)
}

func (sale *SaleTransaction) Fields() []uint8 {
	return sale.entryMap[sale.transData.PosEntryMode]
}

func (sale *SaleTransaction) Name() string {
	return string(sale.transData.TransType)
}
