package ISO8583

import (
	"fmt"
	"strconv"
)

//字段格式
const (
	A      = iota // 字母
	AN            //字幕数字
	ANS           //字幕数字特殊字符
	N             //数字
	Z             //磁道数据
	b             //二进制
	UNUSED        //未使用
)

//长度类型
const (
	FIX    = iota //定长
	VARL          //变长且最大长度小于9
	VARLL         //变长且最大长度小于99
	VARLLL        //变长且最大长度小于999
)

//域字段属性
type Attr struct {
	id      int //域标识
	format  int //域格式
	lenType int //长度类型
	maxLen  int //最大长度
}

//域
type Field struct {
	id    int    //域标识
	value string //值
}

//域属性字典
var fieldAttr = map[int]Attr{
	-2: {-2, N, FIX, 10},
	-1: {-1, N, FIX, 14},
	0:  {0, N, FIX, 4},
	1:  {1, b, FIX, 8},
	2:  {2, N, VARLL, 19},
	3:  {3, N, FIX, 6},

	4: {4, N, FIX, 12},
	5: {5, UNUSED, FIX, 0},
	6: {6, N, FIX, 6},
	7: {7, UNUSED, FIX, 0},
	8: {8, UNUSED, FIX, 0},

	9:  {9, N, FIX, 12},
	10: {10, UNUSED, FIX, 0},
	11: {11, N, FIX, 6},
	12: {12, N, FIX, 6},
	13: {13, N, FIX, 4},

	14: {14, N, FIX, 4},
	15: {15, N, FIX, 4},
	16: {16, UNUSED, FIX, 0},
	17: {17, UNUSED, FIX, 0},
	18: {18, UNUSED, FIX, 0},

	19: {19, UNUSED, UNUSED, 4},
	20: {20, UNUSED, UNUSED, 4},
	21: {21, UNUSED, FIX, 0},
	22: {22, N, FIX, 3},
	23: {23, N, FIX, 3},

	24: {24, N, FIX, 3},
	25: {25, N, FIX, 2},
	26: {26, N, FIX, 2},
	27: {27, UNUSED, FIX, 0},
	28: {28, UNUSED, FIX, 0},

	29: {29, UNUSED, FIX, 4},
	30: {30, UNUSED, FIX, 4},
	31: {31, UNUSED, FIX, 0},
	32: {32, N, VARL, 11},
	33: {33, UNUSED, FIX, 0},

	34: {34, N, UNUSED, 4},
	35: {35, Z, VARLL, 37},
	36: {36, N, VARLLL, 104},
	37: {37, ANS, FIX, 12},
	38: {38, A, FIX, 6},

	39: {39, A, FIX, 2},
	40: {40, UNUSED, FIX, 4},
	41: {41, A, FIX, 8},
	42: {42, A, FIX, 15},
	43: {43, UNUSED, FIX, 0},

	44: {44, A, VARLL, 2},
	45: {45, Z, VARLLL, 4},
	46: {46, UNUSED, FIX, 8},
	47: {47, UNUSED, FIX, 15},
	48: {48, N, VARLLL, 322},

	49: {49, N, FIX, 3},
	50: {50, UNUSED, VARLLL, 4},
	51: {51, N, FIX, 3},
	52: {52, b, FIX, 8},
	53: {53, N, FIX, 16},

	54: {54, A, VARL, 40},
	55: {55, b, VARLLL, 255},
	56: {56, UNUSED, FIX, 8},
	57: {57, UNUSED, FIX, 15},
	58: {58, A, VARLLL, 100},

	59: {59, UNUSED, VARL, 2},
	60: {60, Z, VARLLL, 17},
	61: {61, UNUSED, FIX, 29},
	62: {62, A, VARLLL, 512},
	63: {63, b, VARLLL, 1024},

	64: {64, b, FIX, 8},
}

var fieldRepo = make(map[int]string)

/**
  追加或者修改域的值
  fieldId 域标识
  value 域的值
**/
func SetElement(fieldId int, value string) {
	fieldRepo[fieldId] = value
}

/**
   取域的值
   fieldId 域标识
   返回域的值和错误
**/
func GetElement(fieldId int) (v string, err error) {
	value, OK := fieldRepo[fieldId]

	if OK {
		return value, nil
	} else {
		return value, fmt.Errorf("can't find field:%d\r\n", fieldId)
	}
}

/**
	构建ISO8583报文
	@param fdSets - 构建IS08583报文需要用到的域标识集合
	@retval msg - 构建好的ISO8583报文
	@retval err - 错误
**/
func PrepareISO8583Message(fdSets map[uint8]string) (msg []byte, err error) {
	bitmap := make([]byte, 8)
	vlen := 0
	message := make([]byte, 1024)
	offset := 0
	bitmapOffset := 0

	for id, e := range fdSets {
		attr, ok := fieldAttr[int(id)]

		if !ok {
			err = fmt.Errorf("field attr %d not found\r\n", id)
			return nil, err
		}

		switch attr.format {
		case b:
			vlen = (len(e) + 1) / 2
		default:
			vlen = len(e)
		}

		//		fmt.Printf("[id]:%dvalue:%s\r\n", id, e)
		if vlen > attr.maxLen {
			err = fmt.Errorf("[%d]len invalid, expected:%d, in:%d\r\n", id, attr.maxLen, vlen)
			continue
		}

		switch attr.lenType {
		case FIX:
			if vlen != attr.maxLen {
				err = fmt.Errorf("[%d]len invalid, expected:%d, in:%d\r\n", id, attr.maxLen, vlen)
				continue
			}
		case VARL:
			message[offset] = byte(vlen)
			offset++
		case VARLL:
			message[offset] = byte(((vlen / 10) << 4) | (vlen % 10))
			offset++
		case VARLLL:
			message[offset] = byte(vlen / 100)
			offset++
			message[offset] = byte((((vlen % 100) / 10) << 4) | ((vlen % 100) % 10))
			offset++
		}

		switch attr.format {

		case N:
			if (len(e) % 2) != 0 {
				e = "0" + e
			}

			copy(message[offset:], Base16Decode(e))
			offset += (vlen + 1) / 2
		case Z:
			if (len(e) % 2) != 0 {
				e = e + "F"
			}

			copy(message[offset:], Base16Decode(e))
			offset += (vlen + 1) / 2
		case b:
			copy(message[offset:], Base16Decode(e))
			offset += vlen

		case A, AN, ANS:
			copy(message[offset:], []byte(e))
			offset += vlen

		default:

		}

		if id > 0 {
			bitmap[(id-1)/8] |= 1 << ((8 - id) % 8)
			//fmt.Printf("bitmap:%v", bitmap)
		}

		if id == 0 {
			bitmapOffset = offset
			offset += 8
		}

		//		fmt.Printf("offset:%d\r\n", offset)
		//		fmt.Printf("message now:%s\r\n", Base16Encode(message[0:offset]))
	}

	copy(message[bitmapOffset:bitmapOffset+8], bitmap)
	return message[0:offset], nil
}

/**
   存储参数解析器出来的数据元
   fieldId   域标识
   value  值
   storage 存储位置
   成功返回nil 否则为错误描述
**/
type SaveElement func(fieldId int, value string, storage interface{}) error

/**
	从消息类型开始解析一个构建好的ISO8583消息
	msg 消息类型开始的ISO8583消息
	saveData 存储数据
	storage 存储位置
	成功返回nil 否则为错误描述
**/
func DecodeISO8583Message(msg []byte, saveData SaveElement, storage interface{}) error {
	var ok bool
	var attr Attr
	var id int
	var vlen int
	var value string
	var strlen string

	// fmt.Printf("msg in:%s\r\n", Base16Encode(msg))
	if len(msg) < 10 {
		return fmt.Errorf("invalid message: %s", Base16Encode(msg))
	}

	message := msg[2:]
	bitmap := message[0:8]
	offset := 8

	for i := 0; i < len(bitmap); i++ {
		for j := 7; j >= 0; j-- {
			if (bitmap[i] & (1 << uint(j))) != 0 {
				id = i*8 + (8 - j)
				attr, ok = fieldAttr[id]
				if !ok {
					return fmt.Errorf("field %d attribute not found\r\n", id)
				}

				switch attr.lenType {
				case FIX:
					vlen = attr.maxLen
				case VARL:
					strlen = fmt.Sprintf("%02x", message[offset])
					vlen, _ = strconv.Atoi(strlen)
					offset++
				case VARLL:
					strlen = fmt.Sprintf("%02x", message[offset])
					vlen, _ = strconv.Atoi(strlen)
					offset++
				case VARLLL:
					strlen = fmt.Sprintf("%02x%02x", message[offset], message[offset+1])
					vlen, _ = strconv.Atoi(strlen)
					offset += 2
				}

				// fmt.Printf("id:%d offset:%d-vlen:%d\r\n", id, offset, vlen)

				switch attr.format {
				case N, Z:

					value = Base16Encode(message[offset : offset+((vlen+1)/2)])

					if vlen%2 != 0 {
						saveData(id, value[1:vlen+1], storage)
						fieldRepo[id] = value[1 : vlen+1]
					} else {
						saveData(id, value[:vlen], storage)
						fieldRepo[id] = value[:vlen]
					}
					offset += (vlen + 1) / 2
				case b:

					value = Base16Encode(message[offset : offset+vlen])

					saveData(id, value, storage)
					fieldRepo[id] = value

					offset += vlen
				case A, AN, ANS:
					value = string(message[offset : offset+vlen])
					saveData(id, value, storage)
					fieldRepo[id] = value
					offset += vlen

				default:
				}
			}
		}
	}

	return nil
}
