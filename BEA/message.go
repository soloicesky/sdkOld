package BEA

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"

	"github.com/zhulingbiezhi/sdkOld/ISO8583"
	"github.com/zhulingbiezhi/sdkOld/TLV"

	//	"golang.org/x/crypto/pbkdf2"
)

//加密ISO8583消息
func encryptISO8583Message(msg []byte) []byte {
	keyV1 := ISO8583.Base16Decode("ABCDEF0123456789EEEEEEEEEEEEEEEE")
	keyV2 := ISO8583.Base16Decode("FFFFFFFFFFFFFFFF9876543210FEDCBA")
	var key []byte

	for i := 0; i < len(keyV1); i++ {
		key = append(key, keyV1[i]^keyV2[i])
	}

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, key[:16]...)
	tripleDESKey = append(tripleDESKey, key[:8]...)
	fmt.Println("tripleDESKey ", ISO8583.Base16Encode(tripleDESKey))

	fmt.Printf("msg-----%X")
	encryptedMsg, err := TripleDesEncrypt(msg, tripleDESKey)
	if err != nil {
		fmt.Println("TripleEcbDesEncrypt error :", err.Error())
	}
	return encryptedMsg
}

/**
	根据给定的交易数据和需要打包的位图信息生成ISO8583 报文

**/
func createIISO8583Message(fieldsMap map[uint8]string, config *Config) ([]byte, error) {
	msg, err := ISO8583.PrepareISO8583Message(fieldsMap)
	if err != nil {
		return nil, fmt.Errorf("ISO8583::PrepareISO8583Message error: %s", err.Error())
	}
	fmt.Println("un encode msg: ", ISO8583.Base16Encode(msg))
	encmsg := encryptISO8583Message(msg[10:])
	fmt.Println("encode msg: ", ISO8583.Base16Encode(encmsg))

	dstMsg := make([]byte, 0)
	dstMsg = append(dstMsg, 0x00, 0x00) // len
	dstMsg = append(dstMsg, ISO8583.Base16Decode(config.TPDU)...)
	dstMsg = append(dstMsg, ISO8583.Base16Decode(config.EDS)...)
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
