package BEA

type AuthorizeTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
}

func NewAuthorize(trans *TransactionData, config *Config) *AuthorizeTransaction {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
	return &AuthorizeTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
	}
}

func (auth *AuthorizeTransaction) Valid() error {
	if err := auth.baseValid(); err != nil {
		return err
	}
	return validMatch(auth.transData.Pan,
		auth.transData.Amount,
		auth.transData.TransId,
		auth.transData.CardExpireDate,
		auth.transData.Track2,
	)
}

func (auth *AuthorizeTransaction) SetFields() {
	auth.baseFieldSet()
	auth.set(3, param[auth.transData.TransType].processingCode)
	auth.set(24, param[auth.transData.TransType].nii)
	auth.set(25, param[auth.transData.TransType].posCondictionCode)
}

func (auth *AuthorizeTransaction) Fields() []uint8 {
	return auth.entryMap[auth.transData.PosEntryMode]
}
