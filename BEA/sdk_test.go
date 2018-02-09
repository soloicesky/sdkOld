package BEA

import (
	"fmt"
	"testing"
	//	"github.com/zhulingbiezhi/sdkOld/ISO8583"
	//	"github.com/zhulingbiezhi/sdkOld/TLV"
)

var IccRelatedData = map[string]string{
	"57":   "4761739001010432D22122011631141689",
	"5A":   "4761739001010432",
	"5F2A": "0344",
	"82":   "5C00",
	"84":   "A0000000031010",
	"9A":   "180206",
	"9B":   "E800",
	"9C":   "00",
	"9F02": "000000022500",
	"9F03": "000000000000",
	"9F08": "008D",
	"9F09": "0096",
	"9F1A": "0344",
	"9F1E": "3132333435363738",
	"9F26": "2568BA2CA8FAADA7",
	"9F27": "80",
	"9F33": "E0B8C8",
	"9F34": "1E0300",
	"9F35": "22",
	"9F36": "0001",
	"9F37": "21C8CF43",
	"9F41": "00000130",
	"9F10": "06010A03A00000",
}

var VISIccRelatedData = map[string]string{
	"57":   "4761739001010432D22122011631141689",
	"5A":   "4761739001010432",
	"5F2A": "0344",
	"82":   "5C00",
	"84":   "A0000000031010",
	"95":   "0280008000",
	"9A":   "180206",
	"9B":   "E800",
	"9C":   "00",
	"9F02": "000000106800",
	"9F03": "000000000000",
	"9F08": "008D",
	"9F09": "0096",
	"9F1A": "0344",
	"9F1E": "3132333435363738",
	"9F26": "F205E604A969A4CC",
	"9F27": "80",
	"9F33": "E0B8C8",
	"9F34": "1E0300",
	"9F35": "22",
	"9F36": "0001",
	"9F37": "C1EBF80F",
	"9F41": "00000149",
	"9F10": "06010A03A00000",
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
		TransType:      KindSale,
		Amount:         "00000010000",
		TipAmount:      "1000",
		TransId:        "000149",
		Pan:            "5413330089020029",                  //"5413330089020029D2512201062980790"
		CardExpireDate: "2512",                              //2212
		Track2:         "5413330089020029D2512201062980790", // 4761739001010432D22122011631141689
		PosEntryMode:   SWIPE,
		IccRelatedData: VISIccRelatedData,
	}

	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		// fmt.Println(err)
		t.Errorf(err.Error())
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
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestCapture-------------")
}

func TestRefund(t *testing.T) {
	fmt.Println("---------------Start TestRefund-------------")
	transData := &TransactionData{
		TransType:      KindRefund,
		Amount:         "1000",
		TransId:        "000130",
		Pan:            "5413330089020029",
		CardExpireDate: "2512",
		Track2:         "5413330089020029D2512201062980790",
		PosEntryMode:   INSERT,
		IccRelatedData: IccRelatedData,
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestRefund-------------")
}

func TestVoid(t *testing.T) {
	fmt.Println("---------------Start TestVoid-------------")
	transData := &TransactionData{
		TransType:         KindVoid,
		Amount:            "22500",
		TransId:           "000130",
		Pan:               "5413330089020029",
		CardExpireDate:    "2512",
		Track2:            "5413330089020029D2512201062980790",
		PosEntryMode:      INSERT,
		IccRelatedData:    IccRelatedData,
		OriginalTransType: KindSale,
		AcquireTransID:    "123456789012",
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestVoid-------------")
}

func TestAdjustSale(t *testing.T) {
	fmt.Println("---------------Start TestAdjustSale-------------")
	transData := &TransactionData{
		OriginalTransType: KindSale,
		TransType:         KindAdjustSale,
		Amount:            "109800",
		OriginalAmount:    "106800",
		TipAmount:         "3000",
		TransId:           "000139",
		Pan:               "5413330089020029",
		CardExpireDate:    "2512",
		AcquireTransID:    "180207613032",
		AuthCode:          "005944",
		PosEntryMode:      SWIPE,
		IccRelatedData:    IccRelatedData,
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestAdjustSale-------------")
}

func TestReversal(t *testing.T) {
	fmt.Println("---------------Start TestReversal-------------")
	transData := &TransactionData{
		OriginalTransType: KindSale,
		TransType:         KindReversal,
		Amount:            "106800",
		OriginalAmount:    "106800",
		TransId:           "000139",
		Pan:               "5413330089020029",
		CardExpireDate:    "2512",
		Track2:            "5413330089020029D2512201062980790",
		PosEntryMode:      INSERT,
		IccRelatedData:    IccRelatedData,
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestReversal-------------")
}

func TestSettlement(t *testing.T) {
	fmt.Println("---------------Start TestSettlement-------------")
	transData := &TransactionData{
		TransType: KindSettlmentAfterUpload,
		TransId:   "000139",
		Batchtotals: &BatchTotals{
			CapturedSalesCount:    1,
			CapturedSalesAmount:   100,
			CapturedRefundCount:   0,
			CapturedRefundAmount:  0,
			DebitSalesCount:       0,
			DebitSalesAmount:      0,
			DebitRefundCount:      0,
			DebitRefundAmount:     0,
			AuthorizeSalesCount:   0,
			AuthorizeSalesAmount:  0,
			AuthorizeRefundCount:  0,
			AuthorizeRefundAmount: 0,
		},
		BatchNumber: "000001",
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestSettlement-------------")
}

func TestInitialization(t *testing.T) {
	fmt.Println("---------------Start TestInitialization-------------")
	transData := &TransactionData{
		TransType: KindInitialization,
		TransId:   "000009",
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestInitialization-------------")
}

func TestBatchUpload(t *testing.T) {
	fmt.Println("---------------Start TestBatchUpload-------------")
	transData := &TransactionData{
		TransType:         KindBatchUploadLast,
		Pan:               "5413330089020029",
		CardExpireDate:    "2512",
		TransId:           "000139",
		AcquireTransID:    "180207613032",
		OriginalTransType: KindSale,
		Amount:            "106800",
		AuthCode:          "005944",
		ResponseCode:      "00",
		Track2:            "5413330089020029D2512201062980790",
		TransDate:         "0207",
		TransTime:         "100453",
		PosEntryMode:      SWIPE,
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestBatchUpload-------------")
}

func TestAuthCompletion(t *testing.T) {
	fmt.Println("---------------Start TestBatchUpload-------------")
	transData := &TransactionData{
		TransType:         KindBatchUploadLast,
		Pan:               "5413330089020029",
		CardExpireDate:    "2512",
		TransId:           "000139",
		AcquireTransID:    "180207613032",
		OriginalTransType: KindSale,
		Amount:            "106800",
		AuthCode:          "005944",
		ResponseCode:      "00",
		Track2:            "5413330089020029D2512201062980790",
		TransDate:         "0207",
		TransTime:         "100453",
		PosEntryMode:      SWIPE,
	}
	fmt.Printf("request data: %s", transData.FormJson())
	resp, err := DoRequest(transData, getConfig())
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("response data: %s", resp.FormJson())
	fmt.Println("---------------Success TestBatchUpload-------------")
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
