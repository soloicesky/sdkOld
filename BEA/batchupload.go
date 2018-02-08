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
		INSERT:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 37, 38, 39, 41, 42, 62},
		SWIPE:    {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 39, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 37, 38, 39, 41, 42, 62},
		FALLBACK: {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 39, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 39, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 39, 41, 42, 62},
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
		return nil, fmt.Errorf("unknow transaction type:%s", trans.OriginalTransType)
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
		batchupload.transData.ResponseCode,
		batchupload.transData.AcquireTransID,
		batchupload.transData.TransDate,
		batchupload.transData.TransTime,
	)
}

func (batchupload *BatchUploadTransaction) SetFields() {
	batchupload.baseFieldSet()
	batchupload.set(0, batchupload.messageTypeId)
	batchupload.set(3, batchupload.processingCode)

	var de22 string

	switch batchupload.transData.PosEntryMode {
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

	de22 += "2"
	batchupload.set(22, de22)
	batchupload.set(24, param[batchupload.transData.TransType].nii)
	batchupload.set(25, param[batchupload.transData.TransType].posCondictionCode)
	batchupload.set(37, batchupload.transData.AcquireTransID)
	batchupload.set(38, batchupload.transData.AuthCode)
	batchupload.set(39, string(batchupload.transData.ResponseCode))
	batchupload.set(62, fmt.Sprintf("%012s", batchupload.transData.TransId))
}

func (batchupload *BatchUploadTransaction) Fields() []uint8 {
	return batchupload.entryMap[batchupload.transData.PosEntryMode]
}

func (batchupload *BatchUploadTransaction) Name() string {
	return string(batchupload.transData.TransType)
}
