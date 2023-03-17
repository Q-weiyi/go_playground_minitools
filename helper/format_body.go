package helper

import (
	"encoding/json"
)

func MinifyJSON(jsonStr string) ([]byte, error) {
	var jsonData map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return nil, err
	}

	minifiedJSON, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	return minifiedJSON, nil
}

func ParseJson(jsonStr string, callbackType string) ([]byte, error) {
	var jsonData interface{}
	var err error

	switch callbackType {
	case "directdeposit":
		var data DirectDepositToCallback
		err = json.Unmarshal([]byte(jsonStr), &data)
		jsonData = data
	case "invoicedeposit":
		var data DepositToCallback
		err = json.Unmarshal([]byte(jsonStr), &data)
		jsonData = data
	case "settlement":
		var data SettlementToCallback
		err = json.Unmarshal([]byte(jsonStr), &data)
		jsonData = data
	case "refund":
		var data RefundToCallback
		err = json.Unmarshal([]byte(jsonStr), &data)
		jsonData = data
	case "withdrawal":
		var data WithdrawalToCallback
		err = json.Unmarshal([]byte(jsonStr), &data)
		jsonData = data
	}
	if err != nil {
		return nil, err
	}

	minifiedJSON, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	return minifiedJSON, nil
}
