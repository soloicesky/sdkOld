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
		INSERT:   {0, 3, 11, 24, 41, 42, 60, 62},
		SWIPE:    {0, 3, 11, 24, 41, 42, 60, 62},
		WAVE:     {0, 3, 11, 24, 41, 42, 60, 62},
		FALLBACK: {0, 3, 11, 24, 41, 42, 60, 62},
		MSD:      {0, 3, 11, 24, 41, 42, 60, 62},
		MANUAL:   {0, 3, 11, 24, 41, 42, 60, 62},
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
	return validMatch(settlement.transData.Batchtotals,
		settlement.transData.TransId,
		settlement.transData.BatchNumber,
	)
}

func (settlement *SettlementTransaction) SetFields() {
	settlement.baseFieldSet()
	settlement.set(0, settlement.messageTypeId)
	settlement.set(3, settlement.processingCode)
	settlement.set(24, param[settlement.transData.TransType].nii)
	settlement.set(60, fmt.Sprintf("%06s", settlement.transData.BatchNumber))

	batchTotals := fmt.Sprintf("%03d%012d%03d%012d%03d%012d%03d%012d%03d%012d%03d%012d",
		settlement.transData.Batchtotals.CapturedSalesCount, settlement.transData.Batchtotals.CapturedSalesAmount,
		settlement.transData.Batchtotals.CapturedRefundCount, settlement.transData.Batchtotals.CapturedRefundAmount,
		settlement.transData.Batchtotals.DebitSalesCount, settlement.transData.Batchtotals.DebitSalesAmount,
		settlement.transData.Batchtotals.DebitRefundCount, settlement.transData.Batchtotals.DebitRefundAmount,
		settlement.transData.Batchtotals.AuthorizeSalesCount, settlement.transData.Batchtotals.AuthorizeSalesAmount,
		settlement.transData.Batchtotals.AuthorizeRefundCount, settlement.transData.Batchtotals.AuthorizeRefundAmount)

	settlement.set(62, batchTotals)

}

func (settlement *SettlementTransaction) Fields() []uint8 {
	return []byte{0, 3, 11, 24, 41, 42, 60, 62}
}

func (settlement *SettlementTransaction) Name() string {
	return string(settlement.transData.TransType)
}
