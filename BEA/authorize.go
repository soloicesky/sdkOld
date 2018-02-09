package BEA

import (
	"TLV"
	"encoding/hex"
	"fmt"
)

type AuthorizeTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewAuthorize(trans *TransactionData, config *Config) (*AuthorizeTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
	return &AuthorizeTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0100",
		processingCode: "000000",
	}, nil
}

func (auth *AuthorizeTransaction) Valid() error {
	if err := auth.baseValid(); err != nil {
		return err
	}
	return validMatch(auth.transData.Pan,
		auth.transData.Amount,
		auth.transData.TransId,
		auth.transData.CardExpireDate,
		auth.transData.Track2,
	)
}

func (auth *AuthorizeTransaction) Name() string {
	return "Authorize"
}

func (auth *AuthorizeTransaction) SetFields() {
	auth.baseFieldSet()
	auth.set(0, auth.messageTypeId)
	auth.set(3, auth.processingCode)

	var de22 string

	switch auth.transData.PosEntryMode {
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

	if len(auth.transData.Pin) > 0 {
		de22 += "1"
	} else {
		de22 += "2"
	}

	auth.set(22, de22)
	auth.set(24, param[auth.transData.TransType].nii)
	auth.set(25, param[auth.transData.TransType].posCondictionCode)

	switch auth.transData.PosEntryMode {
	case INSERT:
		fallthrough
	case WAVE:
		var iccData = make(map[string]string)

		for _, tag := range DE55TagList {
			// fmt.Printf("tag:%v\n", tag)
			iccData[tag] = auth.transData.IccRelatedData[tag]
		}

		fmt.Printf("iccdata:%v\r\n", iccData)

		de55 := TLV.BuildConstructTLVMsg(iccData)
		auth.set(55, hex.EncodeToString(de55))
	default:
	}
	auth.set(24, param[auth.transData.TransType].nii)
	auth.set(25, param[auth.transData.TransType].posCondictionCode)
}

func (auth *AuthorizeTransaction) Fields() []uint8 {
	return auth.entryMap[auth.transData.PosEntryMode]
}
