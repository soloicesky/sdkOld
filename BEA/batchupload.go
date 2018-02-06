package BEA

import "fmt"

type BatchUploadTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewBatchUpload(trans *TransactionData, config *Config) (*BatchUploadTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
	trxn := &BatchUploadTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0320",
		processingCode: "000000",
	}

	switch trans.OriginalTransType {
	case KindSale:
		trxn.processingCode = "000000"
	case KindVoidSale:
		trxn.processingCode = "020000"
	case KindRefund:
		trxn.processingCode = "000000"
	case KindVoidRefund:
		trxn.processingCode = "000000"
	case KindPreAuthorize:
		trxn.processingCode = "000000"
	case KindPreAuthCompletion:
		trxn.processingCode = "000000"
	case KindVoidPreAuthorize:
		trxn.processingCode = "000000"
	case KindVoidPreAuthCompletion:
		trxn.processingCode = "000000"
	default:
		return nil, fmt.Errorf("unknow transaction type:", trans.OriginalTransType)
	}

	return trxn, nil
}

func (batchupload *BatchUploadTransaction) Valid() error {
	if err := batchupload.baseValid(); err != nil {
		return err
	}
	return validMatch(batchupload.transData.Pan,
		batchupload.transData.Amount,
		batchupload.transData.TransId,
		batchupload.transData.CardExpireDate,
		batchupload.transData.Track2,
	)
}

func (batchupload *BatchUploadTransaction) SetFields() {
	batchupload.baseFieldSet()
	batchupload.set(3, param[batchupload.transData.TransType].processingCode)
	batchupload.set(24, param[batchupload.transData.TransType].nii)
	batchupload.set(25, param[batchupload.transData.TransType].posCondictionCode)
}

func (batchupload *BatchUploadTransaction) Fields() []uint8 {
	return batchupload.entryMap[batchupload.transData.PosEntryMode]
}

func (batchupload *BatchUploadTransaction) Name() string {
	return string(batchupload.transData.TransType)
}
