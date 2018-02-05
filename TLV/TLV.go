package TLV

import (
	"bindolabs/gateway/services/bea/ISO8583"
	"fmt"
)

/**
	构建TLV数据
	tlvmap TLV 数据集合
**/
func BuildConstructTLVMsg(tlvmap map[string]string) []byte {

	dstMsg := make([]byte, 1)

	for k, v := range tlvmap {
		dstMsg = append(ISO8583.Base16Decode(k))
		if len(v) > 127 {
			dstMsg = append(dstMsg, 0x80)
		}

		dstMsg = append(dstMsg, byte(len(v)))

		if len(v) > 0 {
			dstMsg = append(dstMsg, ISO8583.Base16Decode(v)...)
		}
	}

	return dstMsg
}

/**
	解析已构建的TLV 数据
	tlvMSG 构建的TLV数据
	storrage 存储控件
**/
func ParseConstructTLVMsg(tlvMsg []byte, storage interface{}) {
	iccMap := storage.(map[string]string)
	var tag int
	var length int
	var count int

	for i := 0; i < len(tlvMsg); {
		if tlvMsg[i] == 0x00 || tlvMsg[i] == 0xFF {
			i++
			continue
		}

		tag = int(tlvMsg[i])
		i++

		if (tag & 0x1F) == 0x1F {
			tag = int(tag<<8) + int(tlvMsg[i])
			i++
		}

		fmt.Printf("Tag:%X\r\n", tag)

		if (tlvMsg[i] & 0x80) == 0x80 {
			count = int(tlvMsg[i] & 0x7F)
			length = 0
			i++
		} else {
			length = int(tlvMsg[i])
			count = 0
			i++
		}

		for j := 0; j < count; j++ {
			length = (length << 8) + int(tlvMsg[i])
			i++
		}

		// fmt.Printf("length:%X\r\n", length);
		value := tlvMsg[i : i+length]
		i += length
		stag := fmt.Sprintf("%X", tag)
		// fmt.Printf("value:%s\r\n", ISO8583.Base16Encode(value));
		// fmt.Printf("i=%d-len:%d\r\n", i, len(tlvMsg));
		iccMap[stag] = ISO8583.Base16Encode(value)
	}
}
