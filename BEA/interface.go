package BEA

import (
	"fmt"
	"reflect"
)

type BEAInterface interface {
	Valid() error
	Fields() []byte
	SetFields()
	GetFieldsMap() map[uint8]string
	Name() string
}

type BaseElement struct {
	elementMap map[uint8]string
	transData  *TransactionData
	config     *Config
}

func (e *BaseElement) get(key uint8) string {
	return e.elementMap[key]
}

func (e *BaseElement) set(key uint8, val string) {
	e.elementMap[key] = val
}

func (e *BaseElement) init() {
	e.elementMap = make(map[uint8]string)
}

func (e BaseElement) GetFieldsMap() map[uint8]string {
	return e.elementMap
}

func (e *BaseElement) baseFieldSet() {
	e.init()
	e.set(2, e.transData.Pan)
	e.set(4, fmt.Sprintf("%012s", e.transData.Amount))
	e.set(11, e.transData.TransId)
	e.set(12, e.transData.TransTime)
	e.set(13, e.transData.TransDate)
	e.set(14, e.transData.CardExpireDate)
	e.set(35, e.transData.Track2)
	e.set(41, e.config.TerminalId)
	e.set(42, e.config.MerchantId)
}

func (e *BaseElement) baseValid() error {
	return validMatch(e.transData.TransId,
		e.config.TerminalId,
	)
}

func validMatch(args ...interface{}) error {
	for i := 0; i < len(args); i++ {
		refType := reflect.TypeOf(args[i])
		refValue := reflect.ValueOf(args[i])
		switch refType.Kind() {
		case reflect.String:
			if refValue.String() == "" {
				return fmt.Errorf("%d: the paraments of BEATransaction.string is empty !", i)
			}

		case reflect.Map:
			if len(refValue.MapKeys()) == 0 {
				return fmt.Errorf("%d: the paraments of BEATransaction.map is empty !", i)
			}

		default:
			if refValue.Type().Kind() == reflect.Ptr && refValue.IsNil() {
				return fmt.Errorf("%d: the paraments of BEATransaction.struct is empty !", i)
			}
		}

	}
	return nil
}
