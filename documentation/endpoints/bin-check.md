
**BIN CHECK**
----
Bin numarası ile ilgili ayrıntılı bilgiye [buradan](./../BIN.md) ulaşabilirsiniz.

* **Endpoint : /payment/bin/check**
* **Method:**
  `POST`
* **Data**
```json
{
  "locale": "tr",
  "conversationId": "123456789",
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
        "systemTime": 1519213243242,
        "conversationId": "123456789",
        "binNumber": "554960",
        "cardType": "CREDIT_CARD",
        "cardAssociation": "MASTER_CARD",
        "cardFamily": "Bonus",
        "bankName": "Garanti Bankası",
        "bankCode": 62,
        "commercial": 0
    }
    ```
 
* **Error Response:**

    * **BIN Numarası boş gönderildiğinde:**
    ```JSON
    {
        "status": "failure",
        "errorCode": "5006",
        "errorMessage": "binNumber gönderilmesi zorunludur",
        "locale": "tr",
        "systemTime": 1519210234688,
        "conversationId": "123456789",
        "binNumber": ""
    }
    ```