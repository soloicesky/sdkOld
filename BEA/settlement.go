package BEA

import (
	"fmt"
)

type SettlementTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewSettlement(trans *TransactionData, config *Config) (*SettlementTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}

	trxn := &SettlementTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
	}

	switch trans.TransType {
	case KindSettlment:
		trxn.messageTypeId = "0500"
		trxn.processingCode = "920000"

	case KindSettlmentAfterUpload:
		trxn.messageTypeId = "0500"
		trxn.processingCode = "960000"
	default:
		return nil, fmt.Errorf("unsupport trans type %s", trans.TransType)
	}

	return trxn, nil
}

func (settlement *SettlementTransaction) Valid() error {
	if err := settlement.baseValid(); err != nil {
		return err
	}
	return validMatch(settlement.transData.Pan,
		settlement.transData.Amount,
		settlement.transData.TransId,
		settlement.transData.CardExpireDate,
		settlement.transData.Track2,
	)
}

func (settlement *SettlementTransaction) SetFields() {
	settlement.baseFieldSet()
	settlement.set(3, param[settlement.transData.TransType].processingCode)
	settlement.set(24, param[settlement.transData.TransType].nii)
	settlement.set(25, param[settlement.transData.TransType].posCondictionCode)
}

func (settlement *SettlementTransaction) Fields() []uint8 {
	return settlement.entryMap[settlement.transData.PosEntryMode]
}
