package ssport

import (
	"encoding/json"
	"fmt"
)

// AuthInfo represents auth infomation of data
type AuthInfo struct {
	Token    string
	UserID   int64
	IMEICode string
}

func createAuth(imeicode string) (*AuthInfo, error) {
	url := "https://client4.aipao.me/api/token/QM_Users/LoginSchool?IMEICode=" + imeicode

	rsp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(rsp.Body)

	result := &struct {
		Success bool
		ErrMsg  string
		Data    *AuthInfo
	}{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, fmt.Errorf("verify imeicode error %s: %w", result.ErrMsg, ErrInvalidIMEICode)
	}

	return result.Data, nil
}
