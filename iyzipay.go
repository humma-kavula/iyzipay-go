package Iyzipay

type Iyzipay struct {
	APIKey    string
	APISecret string
	URI       string
}

type ErrorStruct struct {
	Error string
}

func New(APIKey, APISecret, URI string) (*Iyzipay, *ErrorStruct) {
	if APIKey == "" {
		return nil, &ErrorStruct{"APIKey not found."}
	}
	if APISecret == "" {
		return nil, &ErrorStruct{"APISecret not found."}
	}
	if URI == "" {
		return nil, &ErrorStruct{"URI not found."}
	}
	return &Iyzipay{APIKey, APISecret, URI}, nil
}
