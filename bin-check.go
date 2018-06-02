package Iyzipay

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/humma-kavula/iyzipay-go/Utils"
)

type BinCheck struct {
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
	BinNumber      string `json:"binNumber"`
}

type BinCheckSuccess struct {
	Status          string `json:"status"`
	Locale          string `json:"locale"`
	SystemTime      int    `json:"systemTime"`
	ConversationId  string `json:"conversationId"`
	BinNumber       string `json:"binNumber"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily      string `json:"cardFamily"`
	BankName        string `json:"bankName"`
	BankCode        int    `json:"bankCode"`
	Commercial      int    `json:"commercial"`
	ErrorCode       string `json:"errorCode"`
	ErrorMessage    string `json:"errorMessage"`
}

var BinCheckEndpoint string = "/payment/bin/check"

func (i Iyzipay) BinCheck(obj *BinCheck) (*BinCheckSuccess, error) {
	requestBody, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(requestBody))
	req, err := http.NewRequest("POST", i.URI+BinCheckEndpoint, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	RandomString := Utils.RandomString(8)
	pki := Utils.PKIStringify(*obj)

	authorization := Utils.GenerateAuthorizationHeader("IYZWS", i.APIKey, ":", i.APISecret, pki, RandomString)
	req.Header.Add("authorization", authorization)
	req.Header.Add("x-iyzi-rnd", RandomString)
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := &BinCheckSuccess{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
