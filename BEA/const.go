package BEA

import (
	"fmt"
)

type TransactionType string

//交易类型常量
const (
	KindLogon                          TransactionType = "Logon"                          //签到
	KindSale                           TransactionType = "SALE"                           //消费
	KindVoidSale                       TransactionType = "VOIDSALE"                       //消费撤销
	KindRefund                         TransactionType = "REFUND"                         //退货
	KindVoidRefund                     TransactionType = "VOIDREFUND"                     //退货撤销
	KindOfflineSale                    TransactionType = "OFFLINESALE"                    //离线消费
	KindVoidOfflineSale                TransactionType = "VOIDOFFLINESALE"                //离线消费撤销
	KindPreAuthorize                   TransactionType = "PREAUTH"                        //预授权
	KindVoidPreAuthorize               TransactionType = "VOIDPREAUTH"                    //预授权
	KindPreAuthCompletion              TransactionType = "PREAUTHCOMPLETION"              //预授权完成
	KindVoidPreAuthCompletion          TransactionType = "VOIDPREAUTHCOMPLETION"          //预授权完成撤销
	KindSettlment                      TransactionType = "SETTLEMENT"                     //结算
	KindSettlmentAfterUpload           TransactionType = "SETTLEMENTAFTERUPLOAD"          //批上送后结算
	KindBatchUpload                    TransactionType = "BATCHUPLOAD"                    //批上送
	KindBatchUploadLast                TransactionType = "BATCHUPLOADLAST"                //批上送最后一笔
	KindReversal                       TransactionType = "REVERSAL"                       //冲正
	KindOfflineUpload                  TransactionType = "OFFLINEUPLOAD"                  //离线上送
	KindVoidAdjustSale                 TransactionType = "VOIDADJUSTSALE"                 //撤销消费调整
	KindAdjustSale                     TransactionType = "ADJUSTSALE"                     //消费调整
	KindAdjustOfflineSale              TransactionType = "ADJUSTOFFLINESALE"              //离线消费调整
	KindVoidAdjustOfflineSale          TransactionType = "VOIDADJUSTOFFLINESALE"          //离线消费调整撤销
	KindVoidUploadAdjustOfflineSale    TransactionType = "VOIDUPLOADADJUSTOFFLINESALE"    //上送离线消费调整撤销
	KindAdjustUploadOfflineSale        TransactionType = "ADJUSTUPLOADOFFLINESALE"        //上送离线消费调整撤销
	KindVoidAdjustUploadOfflineSale    TransactionType = "VOIDADJUSTUPLOADOFFLINESALE"    //上送离线消费调整撤销
	KindAdjustUploadAdjustSale         TransactionType = "ADJUSTUPLOADADJUSTSALE"         //上送离线消费调整撤销
	KindVoidAdjustUploadAdjustSale     TransactionType = "VOIDADJUSTUPLOADADJUSTSALE"     //上送离线消费调整撤销
	KindAdjustSaleCompletion           TransactionType = "ADJUSTSALECOMPLETION"           //上送离线消费调整撤销
	KindVoidAdjustSaleCompletion       TransactionType = "VOIDADJUSTSALECOMPLETION"       //上送离线消费调整撤销
	KindVoidUploadAdjustSaleCompletion TransactionType = "VOIDUPLOADADJUSTSALECOMPLETION" //上送离线消费调整撤销
	KindAdjustRefund                   TransactionType = "ADJUSTREFUND"                   //调整退货
	KindVoidAdjustRefund               TransactionType = "VOIDADJUSTREFUND"               //调整退货撤销
	KindVoidUploadAdjustRefund         TransactionType = "VOIDUPLOADADJUSTREFUND"         //上送退货调整撤销
	KindVoid                           TransactionType = "VOID"                           //撤销
)

type EntryMode string

//刷卡方式
const (
	INSERT   EntryMode = "CONTACT"        //插IC卡
	SWIPE    EntryMode = "SWIPE"          //刷磁条卡
	WAVE     EntryMode = "CONTACTLESS"    //挥非接卡
	FALLBACK EntryMode = "FALLBACK_SWIPE" //降级
	MSD      EntryMode = "MSD"            //非接卡MSD模式
	MANUAL   EntryMode = "MANUAL"         //手输卡号
)

//终端输入模式码
var posEntryMode = map[EntryMode]string{
	INSERT:   "05", //插卡
	SWIPE:    "90", //刷卡
	WAVE:     "07", //挥卡
	FALLBACK: "80", //降级
	MSD:      "91", //非接磁卡模式
	MANUAL:   "01",
}

//消息参数
type msgParam struct {
	id                string //消息类型
	processingCode    string //处理码
	nii               string //网络信息指示
	posCondictionCode string //POS条件码
}

//消息参数表
var param = map[TransactionType]msgParam{
	KindLogon:                          {"0800", "920000", "028", "00"}, //消费参数
	KindSale:                           {"0200", "000000", "028", "00"}, //消费参数
	KindVoidSale:                       {"0200", "020000", "028", "00"}, //消费撤销参数
	KindRefund:                         {"0200", "200000", "028", "00"}, //退货
	KindVoidRefund:                     {"0200", "220000", "028", "00"}, //撤销退货
	KindOfflineSale:                    {"0220", "000000", "028", "00"}, //离线消费
	KindVoidOfflineSale:                {"0220", "000000", "028", "00"}, //离线消费撤销
	KindPreAuthorize:                   {"0100", "000000", "028", "00"}, //预授权
	KindVoidPreAuthorize:               {"0120", "000000", "028", "00"}, //预授权撤销
	KindPreAuthCompletion:              {"0220", "000000", "028", "00"}, //预授权完成
	KindVoidPreAuthCompletion:          {"0220", "000000", "028", "00"}, //预授权完成撤销
	KindReversal:                       {"0400", "000000", "028", "00"}, //冲正
	KindSettlment:                      {"0500", "920000", "028", "00"}, //结算
	KindSettlmentAfterUpload:           {"0500", "960000", "028", "00"}, //批上送后结算
	KindBatchUpload:                    {"0320", "000000", "028", "00"}, //批上送
	KindBatchUploadLast:                {"0320", "000000", "028", "00"}, //批上送结束
	KindVoidAdjustSale:                 {"0200", "000000", "028", "00"}, //消费调整撤销
	KindAdjustSale:                     {"0220", "000000", "028", "00"}, //消费调整
	KindAdjustOfflineSale:              {"0220", "000000", "028", "00"}, //离线消费调整
	KindVoidAdjustOfflineSale:          {"0220", "000000", "028", "00"}, //离线消费调整撤销
	KindVoidUploadAdjustOfflineSale:    {"0220", "000000", "028", "00"}, //离线消费调整上送撤销
	KindAdjustUploadOfflineSale:        {"0220", "000000", "028", "00"}, //离线消费上送调整
	KindVoidAdjustUploadOfflineSale:    {"0220", "000000", "028", "00"}, //离线消费上送调整撤销
	KindAdjustUploadAdjustSale:         {"0220", "020000", "028", "00"}, //离线调整上送调整
	KindVoidAdjustUploadAdjustSale:     {"0220", "020000", "028", "00"}, //消费调整上送调整撤销
	KindAdjustSaleCompletion:           {"0220", "000000", "028", "00"}, //预授权完成调整
	KindVoidAdjustSaleCompletion:       {"0220", "000000", "028", "00"}, //预授权完成调整撤销
	KindVoidUploadAdjustSaleCompletion: {"0220", "020000", "028", "00"}, //预授权完成调整上送撤销
	KindAdjustRefund:                   {"0220", "220000", "028", "00"}, //退货调整
	KindVoidAdjustRefund:               {"0220", "220000", "028", "00"}, //退货调整撤销
	KindVoidUploadAdjustRefund:         {"0220", "220000", "028", "00"}, //退货调整上送撤销
	KindVoid:                           {"0220", "220000", "028", "00"}, //撤销
}

func getAllEntryModes() []EntryMode {
	var modes []EntryMode
	modes = append(modes, INSERT)
	modes = append(modes, SWIPE)
	modes = append(modes, WAVE)
	modes = append(modes, FALLBACK)
	modes = append(modes, MSD)
	modes = append(modes, MANUAL)
	return modes
}

func getSupportEntryMode(mode EntryMode) error {
	modes := getAllEntryModes()
	for _, v := range modes {
		if mode == v {
			return nil
		}
	}
	return fmt.Errorf("not support pos_entry_mode: %s", mode)
}

var DE55TagList = []string{
	"57",
	"5A",
	"5F2A",
	"82",
	"84",
	"9A",
	"9B",
	"9C",
	"9F02",
	"9F03",
	"9F08",
	"9F09",
	"9F1A",
	"9F1E",
	"9F26",
	"9F27",
	"9F33",
	"9F34",
	"9F35",
	"9F36",
	"9F37",
	"9F41",
	"9F10",
}
