package BEA

import (
	"ISO8583"
	"TLV"
	"fmt"
	"testing"
)

func TestBEA(t *testing.T) {
	var transData TransactionData
	// BEA.Init()

	transData.TransId = "000001"
	transData.Invoice = "000001"
	transData.TransDate = "1221"
	transData.TransTime = "164030"
	// transData.TransType = SALE
	//transData.CurrencyCode = "0344"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	var config Config
	config.TPDU = "7000280000"
	config.EDS = "0003000A00F000"
	config.Host = "192.168.22.188:8081"
	fmt.Printf("%+v\n", transData)
	//transData, _ = Sale(transData, config)
}

func getConfig() Config {
	return Config{
		TPDU:    "7000280000",
		EDS:     "0003000A00F000",
		Host:    "192.168.22.188:8081",
		TimeOut: 30,
	}
}

/*
func TestPreAuthorization(t *testing.T) {
	fmt.Println("------------TestPreAuthorization start-----------------")
	var transData TransactionData
	transData.TransId = "000001"
	// transData.TransType = PREAUTH
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	fmt.Print("request data:\n ", transData.FormJson())
	replyData, err := PreAuthorization(transData, getConfig())

	fmt.Print("%v", err)
	fmt.Print("reply data:\n ", replyData.FormJson())
	fmt.Println("------------TestPreAuthorization end-----------------")
}

func TestPostPreAuthorization(t *testing.T) {
	fmt.Println("------------TestPostPreAuthorization start-----------------")
	var transData TransactionData
	transData.TransId = "000001"
	// transData.TransType = PREAUTHCOMPLETION
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	fmt.Print("request data:\n ", transData.FormJson())
	replyData, err := PreAuthCompletion(transData, getConfig())
	fmt.Print("%v", err)
	fmt.Print("reply data:\n ", replyData.FormJson())
	fmt.Println("------------TestPostPreAuthorization end-----------------")
}

func TestRefund(t *testing.T) {
	fmt.Println("------------TestRefund start-----------------")
	var transData TransactionData
	transData.TransId = "000001"
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	fmt.Print("request data:\n ", transData.FormJson())
	replyData, err := Refund(transData, getConfig())
	fmt.Print("%v", err)
	fmt.Print("reply data:\n ", replyData.FormJson())
	fmt.Println("------------TestRefund end-----------------")
}

func TestSales(t *testing.T) {
	fmt.Println("------------TestSales start-----------------")
	var transData TransactionData
	transData.TransId = "000001"
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	fmt.Print("request data:\n ", transData.FormJson())
	replyData, err := Sale(transData, getConfig())
	fmt.Print("%v", err)
	fmt.Print("reply data:\n ", replyData.FormJson())
	fmt.Println("------------TestSales end-----------------")
}

func TestVoid_PreAuthorization(t *testing.T) {
	fmt.Println("------------TestVoid_PreAuthorization start-----------------")
	var transData TransactionData
	transData.TransId = "000001"
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	fmt.Print("request data:\n ", transData.FormJson())
	replyData, err := VoidPreAuth(transData, getConfig())
	fmt.Print("%v", err)
	fmt.Print("reply data:\n ", replyData.FormJson())
	fmt.Println("------------TestVoid_PreAuthorization end-----------------")
}

func TestVoid_PostPreAuthorization(t *testing.T) {
	fmt.Println("------------TestVoid_PostPreAuthorization start-----------------")
	var transData TransactionData
	transData.TransId = "000001"
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	fmt.Print("request data:\n ", transData.FormJson())
	replyData, err := VoidPreAuthCompletion(transData, getConfig())
	fmt.Print("%v", err)
	fmt.Print("reply data:\n ", replyData.FormJson())
	fmt.Println("------------TestVoid_PostPreAuthorization end-----------------")
}

func TestVoid_Refund(t *testing.T) {
	fmt.Println("------------TestVoidRefund start-----------------")
	var transData TransactionData
	transData.TransId = "000001"
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	fmt.Print("request data:\n ", transData.FormJson())
	replyData, err := VoidRefund(transData, getConfig())
	fmt.Print("%v", err)
	fmt.Print("reply data:\n ", replyData.FormJson())
	fmt.Println("------------TestVoidRefund end-----------------")
}

func TestVoid_Sales(t *testing.T) {
	fmt.Println("------------TestVoid_Sales start-----------------")
	var transData TransactionData
	transData.TransId = "000001"
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	fmt.Print("request data:\n ", transData.FormJson())
	replyData, err := VoidSale(transData, getConfig())
	fmt.Print("%v", err)
	fmt.Print("reply data:\n ", replyData.FormJson())
	fmt.Println("------------TestVoid_Sales end-----------------")
}

func TestReversal(t *testing.T) {
	fmt.Println("------------TestReversal start-----------------")
	var transData TransactionData
	transData.TransId = "000001"
	transData.TransType = SALE
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	fmt.Print("request data:\n ", transData.FormJson())
	replyData, err := Reversal(transData, getConfig())
	fmt.Print("%v", err)
	fmt.Print("reply data:\n ", replyData.FormJson())
	fmt.Println("------------TestReversal end-----------------")
}*/
