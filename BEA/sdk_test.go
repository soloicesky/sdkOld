package BEA

import (
	"fmt"
	"testing"

	//	"github.com/zhulingbiezhi/sdkOld/ISO8583"
	//	"github.com/zhulingbiezhi/sdkOld/TLV"
)

var IccRelatedData = map[string]string{
	"50":   "5649534120435245444954",
	"57":   "4761739001010432D22122011631141689",
	"81":   "0000E484",
	"82":   "5C00",
	"84":   "A0000000031010",
	"87":   "01",
	"90":   "3C96F7658FBC29A202F19146BDE92166B0F6221BBCCB02E326710B9E229D16FAE9AD0C874C0685916E19F0E32693EE201BCE2359509A6D6572F8EC3FC373126B343F9CB8153D61B7EAB2D42DE19D56083185A03DD14C268D40DF0835C55EABFA38ED28BCE42CD0013DA94F800518B753C246EFFBA08FD2029BAD5DFCF0DAF07B7D801C465FFD252C70B92153B330D95DCA2FA1FAAE2D0168A4EA8B475CD805DC32AA964C17BFCD2CD5D0309AB0EA761B",
	"92":   "50DA20DDA8953B693FED84366831BA1EEA97F78F792ACF8CB98FDF0149A7B78FDA1C4967",
	"93":   "52078E99417C94F03F9BBE0D67995AD5244B171FA6B05EEF12B56D0F363EE71808451406F5667426875D18027140228E127258A2011D937539F11770B033E2CD26E47E1FF7FD487688C084A0617D3C189BD164030A6942C5C0D8937E2EAAAF84FFD69AFB550196CD5C935E8F8708156E25074D5B3E6D3365D921B5217E79D3F53666E48C566994D7D69C57A7BC5C770999978BA9315DE223880E3A313426D500D1130B2A474BD9F133A8C922A0452664",
	"94":   "080101001001030018010201",
	"95":   "0280008000",
	"9F37": "595595A7",
	"9F34": "1E0300",
	"9F02": "000000058500",
	"9F08": "008D",
	"9F26": "5538565E1DF18CD4",
	"6F":   "8407A0000000031010A526500B56495341204352454449549F120F4352454449544F20444520564953418701019F110101",
	"A5":   "500B56495341204352454449549F120F4352454449544F20444520564953418701019F110101",
	"5F25": "090701",
	"9F40": "6000F0A001",
	"9F1E": "3132333435363738",
	"9B":   "E800",
	"8F":   "92",
	"5F2A": "0344",
	"9F03": "000000000000",
	"5F24": "221231",
	"9F07": "FF00",
	"9F41": "00000000",
	"9F21": "120017",
	"9F1A": "0344",
	"9F04": "00000000",
	"8C":   "9F02069F03069F1A0295055F2A029A039C019F3704",
	"9F0F": "F040009800",
	"8D":   "8A029F02069F03069F1A0295055F2A029A039C019F3704",
	"4F":   "A0000000031010",
	"9F11": "01",
	"9F45": "DAC5",
	"9F09": "0096",
	"9F1F": "313633313138393030343136303030303030",
	"9F1B": "00003A98",
	"5F28": "0840",
	"9F35": "22",
	"9A":   "180205",
	"9F36": "0001",
	"9C":   "00",
	"9F0E": "0010000000",
	"8E":   "00000000000000001E0302031F00",
	"9F12": "4352454449544F2044452056495341",
	"9F42": "0840",
	"9F33": "E0B8C8",
	"9F10": "06010A03A00000",
	"5A":   "4761739001010432",
	"9F32": "03",
	"9F0D": "F040008800",
	"9F06": "A0000000031010",
	"9F27": "80",
}

func TestAuthorize(t *testing.T) {
	fmt.Println("---------------Start TestAuthorize-------------")
	transData := &TransactionData{
		TransType:      KindPreAuthorize,
		Amount:         "1",
		TransId:        "00001",
		Pan:            "2512",
		CardExpireDate: "",
		Track2:         "5413330056003578D251210100062001602",
		PosEntryMode:   "SWIPE",
		IccRelatedData: IccRelatedData,
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestAuthorize-------------")
}

func TestSales(t *testing.T) {
	fmt.Println("---------------Start TestSales-------------")
	transData := &TransactionData{
		TransType: KindSale,
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestSales-------------")
}

func TestCapture(t *testing.T) {
	fmt.Println("---------------Start TestCapture-------------")
	transData := &TransactionData{
		TransType: KindPreAuthCompletion,
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestCapture-------------")
}

func TestRefund(t *testing.T) {
	fmt.Println("---------------Start TestRefund-------------")
	transData := &TransactionData{
		TransType: KindRefund,
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestRefund-------------")
}

func TestVoid(t *testing.T) {
	fmt.Println("---------------Start TestVoid-------------")
	transData := &TransactionData{
		TransType:      KindVoidSale,
		OriginalAmount: "",
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestVoid-------------")
}

func TestReversal(t *testing.T) {
	fmt.Println("---------------Start TestReversal-------------")
	transData := &TransactionData{
		OriginalTransType: KindSale,
		TransType:         KindReversal,
		OriginalAmount:    "",
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestReversal-------------")
}

func getConfig() *Config {
	return &Config{
		TPDU:       "7000280000",
		EDS:        "0003000A00F000",
		Host:       "bea-uat.bindolabs.com:8081",
		TimeOut:    120,
		MerchantId: "000015204000099",
		TerminalId: "63150001",
	}
}
