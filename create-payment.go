package Iyzipay

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/humma-kavula/iyzipay-go/Utils"
)

type CreatePayment struct {
	Locale          string                        `json:"locale"`
	ConversationId  string                        `json:"conversationId"`
	Price           string                        `json:"price"`
	PaidPrice       string                        `json:"paidPrice"`
	Installment     int                           `json:"installment"`
	PaymentChannel  string                        `json:"paymentChannel"`
	BasketId        string                        `json:"basketId"`
	PaymentGroup    string                        `json:"paymentGroup"`
	PaymentCard     CreatePayment_PaymentCard     `json:"paymentCard"`
	Buyer           CreatePayment_Buyer           `json:"buyer"`
	ShippingAddress CreatePayment_ShippingAddress `json:"shippingAddress"`
	BillingAddress  CreatePayment_BillingAddress  `json:"billingAddress"`
	BasketItems     []CreatePayment_BasketItems   `json:"basketItems"`
	Currency        string                        `json:"currency"`
}

type CreatePayment_PaymentCard struct {
	CardHolderName string `json:"cardHolderName"`
	CardNumber     string `json:"cardNumber"`
	ExpireYear     string `json:"expireYear"`
	ExpireMonth    string `json:"expireMonth"`
	Cvc            string `json:"cvc"`
	RegisterCard   int    `json:"registerCard"`
}

type CreatePayment_Buyer struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	IdentityNumber      string `json:"identityNumber"`
	Email               string `json:"email"`
	GsmNumber           string `json:"gsmNumber"`
	RegistrationDate    string `json:"registrationDate"`
	LastLoginDate       string `json:"lastLoginDate"`
	RegistrationAddress string `json:"registrationAddress"`
	City                string `json:"city"`
	Country             string `json:"country"`
	ZipCode             string `json:"zipCode"`
	Ip                  string `json:"ip"`
}

type CreatePayment_ShippingAddress struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type CreatePayment_BillingAddress struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type CreatePayment_BasketItems struct {
	Id        string `json:"id"`
	Price     string `json:"price"`
	Name      string `json:"name"`
	Category1 string `json:"category1"`
	Category2 string `json:"category2"`
	ItemType  string `json:"itemType"`
}

type CreatePaymentResponse struct {
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
	ItemTransactions             []ItemTransactions `json:"itemTransactions"`
	CardToken                    string             `json:"cardToken"`
	CardUserKey                  string             `json:"cardUserKey"`
	AuthCode                     string             `json:"authCode"`
	Phase                        string             `json:"phase"`
	ErrorCode                    string             `json:"errorCode"`
	ErrorMessage                 string             `json:"errorMessage"`
	ErrorGroup                   string             `json:"errorGroup"`
}

type ItemTransactions struct {
	ItemId                        string          `json:"itemId"`
	PaymentTransactionId          string          `json:"paymentTransactionId"`
	TransactionStatus             int             `json:"transactionStatus"`
	Price                         float32         `json:"price"`
	PaidPrice                     float32         `json:"paidPrice"`
	MerchantCommissionRate        float32         `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float32         `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount      float32         `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float32         `json:"iyziCommissionFee"`
	BlockageRate                  float32         `json:"blockageRate"`
	BlockageRateAmountMerchant    float32         `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float32         `json:"blockageRateAmountSubMerchant"`
	BlockageResolvedDate          string          `json:"blockageResolvedDate"`
	SubMerchantPrice              float32         `json:"subMerchantPrice"`
	SubMerchantPayoutRate         float32         `json:"subMerchantPayoutRate"`
	SubMerchantPayoutAmount       float32         `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float32         `json:"merchantPayoutAmount"`
	ConvertedPayout               ConvertedPayout `json:"convertedPayout"`
}

type ConvertedPayout struct {
	PaidPrice                     float32 `json:"paidPrice"`
	IyziCommissionRateAmount      float32 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float32 `json:"iyziCommissionFee"`
	BlockageRateAmountMerchant    float32 `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float32 `json:"blockageRateAmountSubMerchant"`
	SubMerchantPayoutAmount       float32 `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float32 `json:"merchantPayoutAmount"`
	IyziConversionRate            float32 `json:"iyziConversionRate"`
	IyziConversionRateAmount      float32 `json:"iyziConversionRateAmount"`
	Currency                      string  `json:"currency"`
}

var CreatePaymentEndpoint string = "/payment/auth"

func (i Iyzipay) CreatePayment(obj *CreatePayment) (*CreatePaymentResponse, error) {
	requestBody, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(requestBody))
	req, err := http.NewRequest("POST", i.URI+CreatePaymentEndpoint, payload)
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

	result := &CreatePaymentResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
