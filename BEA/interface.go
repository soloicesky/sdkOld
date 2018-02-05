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

func (e BaseElement) get(key uint8) string {
	return e.elementMap[key]
}

func (e BaseElement) set(key uint8, val string) {
	e.elementMap[key] = val
}

func (e BaseElement) init() {
	e.elementMap = make(map[uint8]string)
}

func (e BaseElement) GetFieldsMap() map[uint8]string {
	return e.elementMap
}

func (e BaseElement) baseFieldSet() {
	e.set(2, e.transData.Pan)
	e.set(4, e.transData.Amount)
	e.set(11, e.transData.TransId)
	e.set(12, e.transData.TransTime)
	e.set(13, e.transData.TransDate)
	e.set(14, e.transData.CardExpireDate)
	e.set(35, e.transData.Track2)
	e.set(37, e.transData.AcquireTransID)
	e.set(38, e.transData.AuthCode)
	e.set(37, e.transData.AcquireTransID)
	e.set(41, e.config.TerminalId)
}

func (e BaseElement) baseValid() error {
	return validMatch(e.transData.Pan,
		e.transData.TransId,
		e.config.TerminalId,
	)
}

func validMatch(args ...interface{}) error {
	for i := 0; i < len(args); i++ {
		refType := reflect.TypeOf(args[i])
		refValue := reflect.ValueOf(args[i])
		if refValue.IsNil() {
			return fmt.Errorf("the paraments of BEATransaction.%s is empty !", refType.Elem().Name())
		}
	}
	return nil
}
