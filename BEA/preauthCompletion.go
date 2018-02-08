package BEA

type CompletionTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewAuthCompletion(trans *TransactionData, config *Config) (*CompletionTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 41, 42, 54, 62},
		SWIPE:    {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 41, 42, 54, 62},
		WAVE:     {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 41, 42, 54, 62},
		FALLBACK: {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 41, 42, 54, 62},
		MSD:      {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 41, 42, 54, 62},
		MANUAL:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 41, 42, 62},
	}

	return &CompletionTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0220",
		processingCode: "000000",
	}, nil
}

func (authcmp *CompletionTransaction) Valid() error {
	if err := authcmp.baseValid(); err != nil {
		return err
	}
	return validMatch(authcmp.transData.Pan,
		authcmp.transData.Amount,
		authcmp.transData.TransId,
		authcmp.transData.CardExpireDate,
		authcmp.transData.AcquireTransID,
		authcmp.transData.AuthCode,
	)
}

func (authcmp *CompletionTransaction) SetFields() {
	authcmp.baseFieldSet()
	authcmp.set(0, authcmp.messageTypeId)
	authcmp.set(3, authcmp.processingCode)

	var de22 string

	switch authcmp.transData.PosEntryMode {
	case INSERT:
		de22 = "05"
	case SWIPE:
		de22 = "90"
	case MANUAL:
		de22 = "01"
	case FALLBACK:
		de22 = "80"
	case WAVE:
		de22 = "70"
	}

	de22 += "2"

	authcmp.set(22, de22)
	authcmp.set(24, param[authcmp.transData.TransType].nii)
	authcmp.set(25, param[authcmp.transData.TransType].posCondictionCode)

	authcmp.set(62, authcmp.transData.TransId)
}

func (authcmp *CompletionTransaction) Fields() []uint8 {
	return authcmp.entryMap[authcmp.transData.PosEntryMode]
}

func (authcmp *CompletionTransaction) Name() string {
	return string(authcmp.transData.TransType)
}
