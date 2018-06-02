package Iyzipay

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/humma-kavula/iyzipay-go/Utils"
)

type FinishSecurePayment struct {
	Locale               string `json:"locale"`
	ConversationId       string `json:"conversationId"`
	PaymentTransactionId string `json:"paymentTransactionId"`
	Price                string `json:"price"`
	Currency             string `json:"currency"`
	Ip                   string `json:"ip"`
}

type FinishSecurePaymentResponse struct {
	Status                       string             `json:"status"`
	Locale                       string             `json:"locale"`
	SystemTime                   int                `json:"systemTime"`
	ConversationId               string             `json:"conversationId"`
	Price                        float32            `json:"price"`
	PaidPrice                    float32            `json:"paidPrice"`
	Installment                  int                `json:"installment"`
	PaymentId                    string             `json:"paymentId"`
	FraudStatus                  int                `json:"fraudStatus"`
	MerchantCommissionRate       float32            `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float32            `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float32            `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float32            `json:"iyziCommissionFee"`
	CardType                     string             `json:"cardType"`
	CardAssociation              string             `json:"cardAssociation"`
	CardFamily                   string             `json:"cardFamily"`
	BinNumber                    string             `json:"binNumber"`
	LastFourDigits               string             `json:"lastFourDigits"`
	BasketId                     string             `json:"basketId"`
	Currency                     string             `json:"currency"`
	Token                        string             `json:"token"`
	PaymentStatus                string             `json:"paymentStatus"`
	ItemTransactions             []ItemTransactions `json:"itemTransactions"` // same as create-payment.go file
	CardToken                    string             `json:"cardToken"`
	CardUserKey                  string             `json:"cardUserKey"`
	AuthCode                     string             `json:"authCode"`
	Phase                        string             `json:"phase"`
	ErrorCode                    string             `json:"errorCode"`
	ErrorMessage                 string             `json:"errorMessage"`
	ErrorGroup                   string             `json:"errorGroup"`
}

var FinishSecurePaymentEndpoint string = "/payment/3dsecure/auth"

func (i Iyzipay) FinishSecurePayment(obj *FinishSecurePayment) (*FinishSecurePaymentResponse, error) {
	requestBody, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(requestBody))
	req, err := http.NewRequest("POST", i.URI+FinishSecurePaymentEndpoint, payload)
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

	result := &FinishSecurePaymentResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
