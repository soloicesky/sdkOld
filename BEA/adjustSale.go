package BEA

import "fmt"

type AdjustSaleTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewAdjustSale(trans *TransactionData, config *Config) (*AdjustSaleTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 54, 60, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 54, 60, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 54, 60, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 54, 60, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 54, 60, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 54, 60, 62},
	}

	return &AdjustSaleTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0220",
		processingCode: "020000",
	}, nil
}

func (adjustSale *AdjustSaleTransaction) Valid() error {
	if err := adjustSale.baseValid(); err != nil {
		return err
	}
	//Transaction = AdjustTips Transaction
	//FromTransaction Sales/Authoriz Transaction
	return validMatch(adjustSale.transData.Pan,
		adjustSale.transData.Amount, //FromTransaction.TotalAmount + Transaction.Amount
		adjustSale.transData.TransId,
		adjustSale.transData.CardExpireDate,
		adjustSale.transData.TipAmount,      //Transaction.Amount
		adjustSale.transData.OriginalAmount, //FromTransacion.TotalAmount
	)
}

func (adjustSale *AdjustSaleTransaction) SetFields() {
	adjustSale.baseFieldSet()
	adjustSale.set(0, adjustSale.messageTypeId)
	adjustSale.set(3, adjustSale.processingCode)

	var de22 string

	switch adjustSale.transData.PosEntryMode {
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

	de22 += "2"

	adjustSale.set(22, de22)
	adjustSale.set(24, param[adjustSale.transData.TransType].nii)
	adjustSale.set(25, param[adjustSale.transData.TransType].posCondictionCode)
	adjustSale.set(37, adjustSale.transData.AcquireTransID)
	adjustSale.set(38, adjustSale.transData.AuthCode)
	adjustSale.set(54, fmt.Sprintf("%012s", adjustSale.transData.TipAmount))
	adjustSale.set(60, fmt.Sprintf("%012s", adjustSale.transData.OriginalAmount))
	adjustSale.set(62, adjustSale.transData.TransId)
}

func (adjustSale *AdjustSaleTransaction) Fields() []uint8 {
	return adjustSale.entryMap[adjustSale.transData.PosEntryMode]
}

func (adjustSale *AdjustSaleTransaction) Name() string {
	return string(adjustSale.transData.TransType)
}
