# PKI (public key infrastructure)

PKI : Iyzico ya gönderilecek olan datanın özel bir sıra ile dizilmiş yorumudur. Authorization header ı oluştururken kullanıyoruz. 

Bir önek ile başlamak gerirse,

```json
{"locale":"tr","conversaMonId":"123456","binNumber":"454671"}
```

Yukarıdaki datayı bir döngüye sokup şu şekle getirmek gerekiyor,

```
[locale=tr,conversaMonId=123456,binNumber=454671]
```

Burada önemli olan şeylerden bir tanesi datanın sıralaması. Eğer "locale" adlı datayı sona koyarsanız (yada ortaya, farketmez...) maalesef geçersiz olacaktır. 

Yukarıdaki örnek sadece BIN sorgulama hizmetinin örneğiydi ve bu maalesef başlı başına PKI sıralamasını anlatmak için yetersiz bir örnek. Örneğin payment durumunda PKI tamamen değişiyor. Örnek vermek gerekirse,

```json
{  
   "locale":"TR",
   "conversationId":"123456789",
   "price":"1.0",
   "paidPrice":"1.5",
   "currency":"TRY",
   "installment":1,
   "paymentChannel":"WEB",
   "basketId":"B67832",
   "paymentGroup":"LISTING",
   "paymentCard":{  
      "cardHolderName":"John Doe",
      "cardNumber":"5400360000000003",
      "expireYear":"2021",
      "expireMonth":"01",
      "cvc":"123",
      "registerCard":0
   },
   "buyer":{  
      "id":"BY789",
      "name":"John",
      "surname":"Doe",
      "identityNumber":"74300864791",
      "email":"email@email.com",
      "gsmNumber":"+905350000000",
      "registrationDate":"2013-04-21 15:12:09",
      "lastLoginDate":"2015-10-05 12:43:35",
      "registrationAddress":"Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
      "city":"Istanbul",
      "country":"Turkey",
      "zipCode":"34732",
      "ip":"85.34.78.112"
   },
   "shippingAddress":{  
      "address":"Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
      "zipCode":"34742",
      "contactName":"Jane Doe",
      "city":"Istanbul",
      "country":"Turkey"
   },
   "billingAddress":{  
      "address":"Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
      "zipCode":"34742",
      "contactName":"Jane Doe",
      "city":"Istanbul",
      "country":"Turkey"
   },
   "basketItems":[  
      {  
         "id":"BI101",
         "price":"0.3",
         "name":"Binocular",
         "category1":"Collectibles",
         "category2":"Accessories",
         "itemType":"PHYSICAL"
      },
      {  
         "id":"BI102",
         "price":"0.5",
         "name":"Game code",
         "category1":"Game",
         "category2":"Online Game Items",
         "itemType":"VIRTUAL"
      },
      {  
         "id":"BI103",
         "price":"0.2",
         "name":"Usb",
         "category1":"Electronics",
         "category2":"Usb / Cable",
         "itemType":"PHYSICAL"
      }
   ]
}
```

Yukarıdaki json datayı göndermek istediğimizde oluşturacağımız hash de PKI değeri olarak ne kullanacağız ? 

Her bir objenin içerisindeki değeri alıp ilk örnekte verdiğimiz gibi "[" ve "]" arasına koymak gerekiyor. Ancak birden fazla objenin iç içe olduğu durumda ve o objenin de içerisinde bir array olduğu durumda işler biraz farklı. İsterseniz öncelikle olması gereken PKI yi kontrol edelim.

```
[locale=TR,conversationId=123456789,price=1.0,paidPrice=1.5,installment=1,paymentChannel=WEB,basketId=B67832,paymentGroup=LISTING,paymentCard=[cardHolderName=John Doe,cardNumber=5400360000000003,expireYear=2021,expireMonth=01,cvc=123,registerCard=0],buyer=[id=BY789,name=John,surname=Doe,identityNumber=74300864791,email=email@email.com,gsmNumber=+905350000000,registrationDate=2013-04-21 15:12:09,lastLoginDate=2015-10-05 12:43:35,registrationAddress=Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1,city=Istanbul,country=Turkey,zipCode=34732,ip=85.34.78.112],shippingAddress=[address=Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1,zipCode=34742,contactName=Jane Doe,city=Istanbul,country=Turkey],billingAddress=[address=Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1,zipCode=34742,contactName=Jane Doe,city=Istanbul,country=Turkey],basketItems=[[id=BI101,price=0.3,name=Binocular,category1=Collectibles,category2=Accessories,itemType=PHYSICAL], [id=BI102,price=0.5,name=Game code,category1=Game,category2=Online Game Items,itemType=VIRTUAL], [id=BI103,price=0.2,name=Usb,category1=Electronics,category2=Usb / Cable,itemType=PHYSICAL]],currency=TRY]
```

İlk dikkat etmeniz gereken şey "currency" değişkeninin yeri en altta olması. Açıkçası sebebini bilmiyorum. Yaptığım testlerde fark ettim ve çalıştı. 

Diğer dikkat etmeniz gereken yer basketItems değişkenin içerisinde yer alan değişkenlerin arasındaki boşluk. Diğer hiçbir yerde boşluk olmadığı için dikkatten kaçması kolay. Bunun sebebi iyziconun Java dilindeki “ToStringBuilder” fonksiyonunun oluşturduğu formatı temel alması. Siz buna çok dikkat edin.

PKI değerinin hesaplanması sırasında hangi sıra ile yerleştirmemiz gerektiği ve bu sıralamada nerelere dikkat etmeniz gerektiği gibi bilgiler şu an bu yazıyı yazdığım tarihlerde maalesef yok. O yüzden bende sizlere test ederek bulduğum bilgileri sunuyorum.

Kıasacası eğer imza hatası alıyorsanız büyük ihtimal PKI değerinin sıralamasını yanlış sıra ile girmiş olabilirsiniz.