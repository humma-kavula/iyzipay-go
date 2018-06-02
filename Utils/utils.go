package Utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type ApiMethods struct {
	RETRIEVE string
	CREATE   string
	DELETE   string
	UPDATE   string
}

var APIMethod = &ApiMethods{"retrieve", "create", "delete", "update"}

// GenerateAuthorizationHeader : Creting authorization header for request...
func GenerateAuthorizationHeader(iyziWsHeaderName, apiKey, separator, secretKey, body, randomString string) string {
	return iyziWsHeaderName + " " + apiKey + separator + GenerateHash(apiKey, randomString, secretKey, body)
}

// GenerateHash : Generating hash
func GenerateHash(apiKey, randomString, secretKey, body string) string {
	hasher := sha1.New()
	hasher.Write([]byte(apiKey + randomString + secretKey + body))
	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

// RandomString : Generating romdom string
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandomString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// PKI string generator...
func PKIStringify(v interface{}) (res string) {
	rv := reflect.ValueOf(v)
	num := rv.NumField()
	for i := 0; i < num; i++ {
		fv := rv.Field(i)
		st := rv.Type().Field(i)
		jsonName, err := st.Tag.Lookup("json")
		if err != true {
			res += st.Name + "="
		} else {
			res += jsonName + "="
		}
		switch fv.Kind() {
		case reflect.String:
			res += fv.String()
		case reflect.Int:
			res += fmt.Sprint(fv.Int())
		case reflect.Struct:
			res += PKIStringify(fv.Interface())
		case reflect.Slice:
			res += "["
			for k := 0; k < fv.Len(); k++ {
				e := fv.Index(k)
				res += PKIStringify(e.Interface())
				if k != fv.Len()-1 {
					res += ", "
				}
			}
			res += "]"
		}
		if i != num-1 {
			res += ","
		}
	}
	return "[" + res + "]"
}
