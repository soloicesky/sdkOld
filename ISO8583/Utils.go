package ISO8583

import(
_"fmt"
)
/**
	将byte字节进行base16编码转化为十六进制的可见字符串
**/
func Base16Encode( message []byte) string {
	chars := []byte{'0','1','2','3','4','5','6','7','8','9','A', 'B','C','D','E','F'}
	encodeMsg := make([]byte, len(message)*2)
	i:=0

	for _,e := range message {
		encodeMsg[i] = chars[(e&0xF0)>>4]
		encodeMsg[i+1] = chars[e&0x0F]
		i+=2
	}

	return string(encodeMsg)
}

func charToDigit(ch byte) byte {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' :
			return (ch - '0')
	case 'A','a':
		return 0x0A
	case 'B', 'b':
		return 0x0B
	case 'C','c':
		return 0x0C
	case 'D','d':
		return 0x0D
	case 'E','e':
		return 0x0E
	case 'F', 'f':
		return 0x0F
	default:

		return 0x00
	}

	return 0x00
}

/**
	将16进制字符串解码转化为字节数组
**/
func Base16Decode(message string) []byte {
	decodeMsg := make([]byte, (len(message) + 1) / 2)

	for i := 0; i < len(message); {
		decodeMsg[i/2] = charToDigit(message[i])<< 4 | charToDigit(message[i+1])
		i+=2
	}

	return decodeMsg
}

/**
	判断字符串是否为空

**/
func StringIsEmpty(str string)(empty bool) {
	empty = false

	if len(str)<= 0 {
		empty = true
	}

	return empty
}
