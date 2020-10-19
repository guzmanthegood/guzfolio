package model

import (
	"fmt"
	"io"
	"strconv"
)

type Currency struct {
	ID   uint         `json:"id"`
	Code string       `json:"code" gorm:"unique"`
	Name string       `json:"name"`
	Type CurrencyType `json:"type"`
}

type CreateCurrencyInput struct {
	Code string       `json:"code"`
	Name string       `json:"name"`
	Type CurrencyType `json:"type"`
}

type CurrencyType string

const (
	CurrencyTypeFiat   CurrencyType = "FIAT"
	CurrencyTypeCrypto CurrencyType = "CRYPTO"
)

var AllCurrencyType = []CurrencyType{
	CurrencyTypeFiat,
	CurrencyTypeCrypto,
}

func (e CurrencyType) IsValid() bool {
	switch e {
	case CurrencyTypeFiat, CurrencyTypeCrypto:
		return true
	}
	return false
}

func (e CurrencyType) String() string {
	return string(e)
}

func (e *CurrencyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CurrencyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CurrencyType", str)
	}
	return nil
}

func (e CurrencyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}