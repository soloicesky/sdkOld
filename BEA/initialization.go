package BEA

type InitializeTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewInitialization(trans *TransactionData, config *Config) (*InitializeTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}

	return &InitializeTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0800",
		processingCode: "930000",
	}, nil
}

func (initialization *InitializeTransaction) Valid() error {
	if err := initialization.baseValid(); err != nil {
		return err
	}
	return validMatch(initialization.transData.TransId)
}

func (initialization *InitializeTransaction) SetFields() {
	initialization.baseFieldSet()
	initialization.set(0, initialization.messageTypeId)
	initialization.set(3, initialization.processingCode)
	initialization.set(24, param[initialization.transData.TransType].nii)
	initialization.set(62, initialization.transData.TransId)
}

func (initialization *InitializeTransaction) Fields() []uint8 {
	return []byte{0, 3, 11, 24, 41, 42, 62}
}

func (initialization *InitializeTransaction) Name() string {
	return string(initialization.transData.TransType)
}
