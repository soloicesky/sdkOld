package BEA

import (
	"ISO8583"
	"TLV"
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
	//	"golang.org/x/crypto/pbkdf2"
)

//加密ISO8583消息
func encryptISO8583Message(msg []byte) []byte {
	keyV1, _ := hex.DecodeString("ABCDEF0123456789EEEEEEEEEEEEEEEE")
	keyV2, _ := hex.DecodeString("FFFFFFFFFFFFFFFF9876543210FEDCBA")
	var key []byte

	for i := 0; i < len(keyV1); i++ {
		key = append(key, keyV1[i]^keyV2[i])
	}

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, key[:16]...)
	tripleDESKey = append(tripleDESKey, key[:8]...)
	fmt.Println("tripleDESKey ", hex.EncodeToString(tripleDESKey))

	// fmt.Printf("msg-----%X")
	encryptedMsg, err := TripleDesEncrypt(msg, tripleDESKey)

	if err != nil {
		fmt.Println("TripleEcbDesEncrypt error :", err.Error())
	}

	return encryptedMsg
}

/**
	根据给定的交易数据和需要打包的位图信息生成ISO8583 报文

**/
func createIISO8583Message(transData *TransactionData, fields []byte, config *Config) ([]byte, error) {
	ISO8583.SetElement(0, param[transData.TransType].id)
	ISO8583.SetElement(2, transData.Pan)
	if transData.TransType == KindReversal {
		ISO8583.SetElement(3, param[transData.OriginalTransType].processingCode)
		ISO8583.SetElement(24, param[transData.OriginalTransType].nii)
		ISO8583.SetElement(25, param[transData.OriginalTransType].posCondictionCode)
	} else {
		ISO8583.SetElement(3, param[transData.TransType].processingCode)
		ISO8583.SetElement(24, param[transData.TransType].nii)
		ISO8583.SetElement(25, param[transData.TransType].posCondictionCode)
	}

	ISO8583.SetElement(35, transData.Track2)
	ISO8583.SetElement(4, fmt.Sprintf("%012s", transData.Amount))
	ISO8583.SetElement(11, fmt.Sprintf("%06s", transData.TransId))
	ISO8583.SetElement(14, transData.CardExpireDate)

	if ISO8583.StringIsEmpty(transData.Pin) {
		ISO8583.SetElement(22, posEntryMode[transData.PosEntryMode]+"2")
	} else {
		ISO8583.SetElement(22, posEntryMode[transData.PosEntryMode]+"1")
	}

	if !ISO8583.StringIsEmpty(transData.PanSeqNo) {
		ISO8583.SetElement(23, fmt.Sprintf("%04s", transData.PanSeqNo))
	}

	if !ISO8583.StringIsEmpty(transData.AcquireTransID) {
		ISO8583.SetElement(37, transData.AcquireTransID)
	}

	if !ISO8583.StringIsEmpty(config.TerminalId) {
		ISO8583.SetElement(41, config.TerminalId)
	}

	if !ISO8583.StringIsEmpty(config.MerchantId) {
		ISO8583.SetElement(42, config.MerchantId)
	}

	DE55 := TLV.BuildConstructTLVMsg(transData.IccRelatedData)
	ISO8583.SetElement(55, hex.EncodeToString(DE55))

	if !ISO8583.StringIsEmpty(transData.OriginalAmount) {
		ISO8583.SetElement(60, fmt.Sprintf("%012s", transData.OriginalAmount))
	}

	if !ISO8583.StringIsEmpty(transData.Invoice) {
		ISO8583.SetElement(62, transData.Invoice)
	}

	batchTotal := fmt.Sprintf("%03d%012d%03d%012d%03d%012d%03d%012d%03d%012d%03d%012d",
		transData.Batchtotals.CapturedSalesCount, transData.Batchtotals.CapturedSalesAmount,
		transData.Batchtotals.CapturedRefundCount, transData.Batchtotals.CapturedRefundAmount,
		transData.Batchtotals.DebitSalesCount, transData.Batchtotals.DebitSalesAmount,
		transData.Batchtotals.DebitRefundCount, transData.Batchtotals.DebitRefundAmount,
		transData.Batchtotals.AuthorizeSalesCount, transData.Batchtotals.AuthorizeSalesAmount,
		transData.Batchtotals.AuthorizeRefundCount, transData.Batchtotals.AuthorizeRefundAmount)

	ISO8583.SetElement(63, batchTotal)
	msg, err := ISO8583.PrepareISO8583Message(fields)
	if err != nil {
		return nil, fmt.Errorf("ISO8583::PrepareISO8583Message error: %s", err.Error())
	}
	fmt.Println("un encode msg: ", hex.EncodeToString(msg))
	encmsg := encryptISO8583Message(msg[10:])
	fmt.Println("encode msg: ", hex.EncodeToString(encmsg))

	dstMsg := make([]byte, 0)
	dstMsg = append(dstMsg, 0x00, 0x00) // len
	tpdu, _ := hex.DecodeString(config.TPDU)
	dstMsg = append(dstMsg, tpdu...)
	eds, _ := hex.DecodeString(config.EDS)
	dstMsg = append(dstMsg, eds...)
	dstMsg = append(dstMsg, msg[:10]...)
	dstMsg = append(dstMsg, encmsg...)
	dstMsg[2+5+4] = byte(len(encmsg) >> 8)
	dstMsg[2+5+5] = byte(len(encmsg) & 0x00FF)

	dstMsg[0] = byte((len(dstMsg) - 2) >> 8)
	dstMsg[1] = byte((len(dstMsg) - 2) & 0x00FF)

	return dstMsg, nil
}

// 3DES加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	// origData = PKCS5Padding(origData, block.BlockSize())
	origData = ZeroPadding(origData, block.BlockSize())
	IV := make([]byte, 8)
	blockMode := cipher.NewCBCEncrypter(block, IV)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
