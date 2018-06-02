package Iyzipay

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/humma-kavula/iyzipay-go/Utils"
)

type InitializeSecurePayment struct {
	Locale          string                                  `json:"locale"`
	ConversationId  string                                  `json:"conversationId"`
	Price           string                                  `json:"price"`
	PaidPrice       string                                  `json:"paidPrice"`
	Installment     int                                     `json:"installment"`
	PaymentChannel  string                                  `json:"paymentChannel"`
	BasketId        string                                  `json:"basketId"`
	PaymentGroup    string                                  `json:"paymentGroup"`
	PaymentCard     InitializeSecurePayment_PaymentCard     `json:"paymentCard"`
	Buyer           InitializeSecurePayment_Buyer           `json:"buyer"`
	ShippingAddress InitializeSecurePayment_ShippingAddress `json:"shippingAddress"`
	BillingAddress  InitializeSecurePayment_BillingAddress  `json:"billingAddress"`
	BasketItems     []InitializeSecurePayment_BasketItems   `json:"basketItems"`
	Currency        string                                  `json:"currency"`
	SoftDescriptor  string                                  `json:"softDescriptor"`
	CallbackUrl     string                                  `json:"callbackUrl"`
}

type InitializeSecurePayment_PaymentCard struct {
	CardHolderName string `json:"cardHolderName"`
	CardNumber     string `json:"cardNumber"`
	ExpireYear     string `json:"expireYear"`
	ExpireMonth    string `json:"expireMonth"`
	Cvc            string `json:"cvc"`
	RegisterCard   int    `json:"registerCard"`
}

type InitializeSecurePayment_Buyer struct {
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

type InitializeSecurePayment_ShippingAddress struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type InitializeSecurePayment_BillingAddress struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type InitializeSecurePayment_BasketItems struct {
	Id        string `json:"id"`
	Price     string `json:"price"`
	Name      string `json:"name"`
	Category1 string `json:"category1"`
	Category2 string `json:"category2"`
	ItemType  string `json:"itemType"`
}

type InitializeSecurePaymentResponse struct {
	Status             string `json:"status"`
	Locale             string `json:"locale"`
	SystemTime         int    `json:"systemTime"`
	ConversationId     string `json:"conversationId"`
	ThreeDSHtmlContent string `json:"threeDSHtmlContent"`
	ErrorCode          string `json:"errorCode"`
	ErrorMessage       string `json:"errorMessage"`
}

var InitializeSecurePaymentEndpoint string = "/payment/3dsecure/initialize"

func (i Iyzipay) InitializeSecurePayment(obj *InitializeSecurePayment) (*InitializeSecurePaymentResponse, error) {
	requestBody, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(requestBody))
	req, err := http.NewRequest("POST", i.URI+InitializeSecurePaymentEndpoint, payload)
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

	result := &InitializeSecurePaymentResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
