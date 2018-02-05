package BEA

type BEACode string

func (err BEACode) String() string {
	return errorMap[err]
}

const (
	APPROVED                       BEACode = "00" //交易批准
	CALL_ISSUER                    BEACode = "01" //请联系放卡行
	REFFERAL                       BEACode = "02" //参考交易，请联系发卡行
	INVALID_MERCHANT               BEACode = "03" //无效商户
	PICK_UP_CARD                   BEACode = "04" //没收无效卡
	DO_NOT_HONOUR                  BEACode = "05" //不允许，不接受
	INVALID_TRANSACTION            BEACode = "12" //无效交易
	INVALID_AMOUNT                 BEACode = "13" //无效金额
	INVALID_CARD_NUMBER            BEACode = "14" //无效卡号
	REENTER_TRANSACTION            BEACode = "19" //重复交易
	NO_TRANSACTIONS                BEACode = "21" //找不到交易
	UNABLE_LOCATE_RECORD           BEACode = "25" //
	FORMAT_ERROR                   BEACode = "30"
	BANK_NORT_SUPPORTED            BEACode = "31"
	LOST_CARD                      BEACode = "41"
	STOLEN_CARD                    BEACode = "43"
	INSUFFICIENT_FUNDS             BEACode = "51"
	NO_CHEQUING_ACCOUNT            BEACode = "52"
	NO_SAVINGS_ACCOUNT             BEACode = "53"
	EXPIRED_CARD                   BEACode = "54"
	INCORRECT_PIN                  BEACode = "55"
	NO_CARD_RECORD                 BEACode = "56"
	NOT_PERMITTED                  BEACode = "58" //交易批准
	EXCEED_WITHDRAWAL_AMOUNT_LIMIT BEACode = "61"
	SECURITY_VIOLATION             BEACode = "63"
	PIN_TRIES_EXCEED               BEACode = "75"
	INVALID_PRODUCT_CODE           BEACode = "76"
	RECONCILE_ERROR                BEACode = "77"
	TRANS_NOT_FOUND                BEACode = "78"
	BATCH_ALREADY_OPEN             BEACode = "79"
	BATCH_NUMBER_NOT_FOUND         BEACode = "80"
	BATCH_NOT_FOUND                BEACode = "85"
	BAT_TERMINAL_ID                BEACode = "89"
	ISSUER_OR_SWITCH_INOPERATIVE   BEACode = "91"
	DUPLICATE_TRANSMISSION         BEACode = "94"
	BATCH_TRANSFER                 BEACode = "95"
	SYSTEM_MALFUNCTION             BEACode = "96"

	//Locally Generated Error Messages
	LE_LOST_CARRIER                  BEACode = "LC"
	LE_COMM_ERROR                    BEACode = "CE"
	LE_INVALID_DOWNLOADLINE_LOAD     BEACode = "ID"
	LE_INVALID_AMOUNT                BEACode = "IA"
	LE_INVALID_MESSAGE_TYPE          BEACode = "IR"
	LE_INVALID_HOST_SEQUENCE_NUMBER  BEACode = "IS"
	LE_INVALID_MAC                   BEACode = "IM"
	LE_NO_ERPLY_TIMEOUT              BEACode = "TO"
	LE_ADIVICE_REVERSAL_NOT_APPROVED BEACode = "ND"
)

var errorMap = map[BEACode]string{
	APPROVED:                       "Approved",
	CALL_ISSUER:                    "Refer to card issuer",
	REFFERAL:                       "Refer to card Issuer\"s special conditions", //参考交易，请联系发卡行
	INVALID_MERCHANT:               "Invalid merchant",                           //无效商户
	PICK_UP_CARD:                   "Pick-up",                                    //没收无效卡
	DO_NOT_HONOUR:                  "Do not honour",                              //不允许，不接受
	INVALID_TRANSACTION:            "Invalid transaction",                        //无效交易
	INVALID_AMOUNT:                 "Invalid amount",                             //无效金额
	INVALID_CARD_NUMBER:            "Invalid card number",                        //无效卡号
	REENTER_TRANSACTION:            "Re-enter transaction",                       //重复交易
	NO_TRANSACTIONS:                "No Transactions",                            //找不到交易
	UNABLE_LOCATE_RECORD:           "Unable to locate record on file",            //
	FORMAT_ERROR:                   "Format error",
	BANK_NORT_SUPPORTED:            "Bank not supported by switch",
	LOST_CARD:                      "Lost card",
	STOLEN_CARD:                    "Stolen card, pick up",
	INSUFFICIENT_FUNDS:             "Not sufficient funds",
	NO_CHEQUING_ACCOUNT:            "No chequing account",
	NO_SAVINGS_ACCOUNT:             "No savings account",
	EXPIRED_CARD:                   "Expired card",
	INCORRECT_PIN:                  "Incorrect PIN",
	NO_CARD_RECORD:                 "No card record",
	NOT_PERMITTED:                  "Transaction not permitted to terminal", //交易批准
	EXCEED_WITHDRAWAL_AMOUNT_LIMIT: "Exceeds withdrawal amount limit",
	SECURITY_VIOLATION:             "Security violation",
	PIN_TRIES_EXCEED:               "Allowable number of PIN tries exceeded",
	INVALID_PRODUCT_CODE:           "Invalid product code",
	RECONCILE_ERROR:                "Reconcile error (or host text if sent)",
	TRANS_NOT_FOUND:                "Trans. number not found",
	BATCH_ALREADY_OPEN:             "Batch already open",
	BATCH_NUMBER_NOT_FOUND:         "Batch number not found",
	BATCH_NOT_FOUND:                "Batch not found",
	BAT_TERMINAL_ID:                "Bad Terminal ID",
	ISSUER_OR_SWITCH_INOPERATIVE:   "Issuer or switch inoperative",
	DUPLICATE_TRANSMISSION:         "Duplicate transmission",
	BATCH_TRANSFER:                 "Reconcile error. Batch upload started",
	SYSTEM_MALFUNCTION:             "System malfunction",

	//LocallyGeneratedErrorMessages
	LE_LOST_CARRIER:                  "Lost carrier",
	LE_COMM_ERROR:                    "Communications error",
	LE_INVALID_DOWNLOADLINE_LOAD:     "Invalid downline load",
	LE_INVALID_AMOUNT:                "Invalid amount",
	LE_INVALID_MESSAGE_TYPE:          "Invalid message type",
	LE_INVALID_HOST_SEQUENCE_NUMBER:  "nvalid host sequence number",
	LE_INVALID_MAC:                   "Invalid MAC",
	LE_NO_ERPLY_TIMEOUT:              "No reply timeout",
	LE_ADIVICE_REVERSAL_NOT_APPROVED: "Advice / Reversal transactions not approved",

	//bindo generated error code
	BINDO_CONN_ERR: "fail to connect to host",   //连接失败
	BINDO_SEND_ERR: "fail to send data to host", //发送失败
	BINDO_RECV_ERR: "fail to recv from host",    //接收失败
	BINDO_COMM_ERR: "bindo unkown error",        //未知错误
}
