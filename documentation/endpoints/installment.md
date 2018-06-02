
**Installment**
----

Yapılabilinecek taksit bilgilerini döndürür.

* **Endpoint : /payment/iyzipos/installment**
* **Method:**`POST`
* **Data**
```json
{
  "locale": "tr",
  "conversationId": "123456789",
  "price":"100.0",
  "binNumber":"554960"
}
```
* **Success Response:**
  * **Code:** 200 <br />
    **Content:**
    ```json
    {
        "status": "success",
        "locale": "tr",
        "systemTime": 1519223095946,
        "conversationId": "123456789",
        "installmentDetails": [
            {
                "binNumber": "554960",
                "price": 100,
                "cardType": "CREDIT_CARD",
                "cardAssociation": "MASTER_CARD",
                "cardFamilyName": "Bonus",
                "force3ds": 0,
                "bankCode": 62,
                "bankName": "Garanti Bankası",
                "forceCvc": 0,
                "commercial": 0,
                "installmentPrices": [
                    {
                        "installmentPrice": 100,
                        "totalPrice": 100,
                        "installmentNumber": 1
                    },
                    {
                        "installmentPrice": 50.44,
                        "totalPrice": 100.87,
                        "installmentNumber": 2
                    },
                    {
                        "installmentPrice": 33.88,
                        "totalPrice": 101.64,
                        "installmentNumber": 3
                    },
                    {
                        "installmentPrice": 17.23,
                        "totalPrice": 103.4,
                        "installmentNumber": 6
                    },
                    {
                        "installmentPrice": 11.65,
                        "totalPrice": 104.86,
                        "installmentNumber": 9
                    }
                ]
            }
        ]
    }
    ```
* **Error Response:**
```json
{
    "status": "failure",
    "errorCode": "1000",
    "errorMessage": "Geçersiz imza",
    "locale": "tr",
    "systemTime": 1519223361805,
    "conversationId": "123456789"
}
```