package Iyzipay

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/humma-kavula/iyzipay-go/Utils"
)

type InstallmentCheck struct {
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
	BinNumber      string `json:"binNumber"`
	Price          string `json:"price"`
}

type InstallmentCheckSuccess struct {
	Status             string                     `json:"status"`
	Locale             string                     `json:"locale"`
	SystemTime         int                        `json:"systemTime"`
	ConversationId     string                     `json:"conversationId"`
	InstallmentDetails []InstallmentDetailsStruct `json:"installmentDetails"`
	ErrorCode          string                     `json:"errorCode"`
	ErrorMessage       string                     `json:"errorMessage"`
}

type InstallmentDetailsStruct struct {
	BinNumber         string              `json:"binNumber"`
	Price             float32             `json:"price"`
	CardType          string              `json:"cardType"`
	CardAssociation   string              `json:"cardAssociation"`
	CardFamilyName    string              `json:"cardFamilyName"`
	Force3ds          int                 `json:"force3ds"`
	BankCode          int                 `json:"bankCode"`
	BankName          string              `json:"bankName"`
	ForceCvc          int                 `json:"forceCvc"`
	Commercial        int                 `json:"commercial"`
	InstallmentPrices []InstallmentPrices `json:"installmentPrices"`
}

type InstallmentPrices struct {
	InstallmentPrice  float32 `json:"installmentPrice"`
	TotalPrice        float32 `json:"totalPrice"`
	InstallmentNumber float32 `json:"installmentNumber"`
}

var InstallmentCheckEndpoint string = "/payment/iyzipos/installment"

func (i Iyzipay) InstallmentCheck(obj *InstallmentCheck) (*InstallmentCheckSuccess, error) {
	requestBody, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(requestBody))
	req, err := http.NewRequest("POST", i.URI+InstallmentCheckEndpoint, payload)
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
	result := &InstallmentCheckSuccess{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
