package core

import (
	"encoding/json"
	"net/http"
)

const (
	SELLER_SERVICE = "http://seller-core-api.seller-service"

	USER_SERVICE = "http://iam-oauth-api-app.iam-service"

	SRM_SERVICE = "http://supplier-api-v2-api.supplier-management"
)

func GetUserInfoByToken(token string) (interface{}, error) {
	url := SELLER_SERVICE + "/userinfo"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}

	var text []byte
	_, err = resp.Body.Read(text)
	if err != nil {
		return nil, err
	}

	var data interface{}
	if err := json.Unmarshal(text, data); err != nil {
		return nil, err
	}
	return data, nil
}
