# BIN (Base Identification Number)

Kredi/Banka kartının ilk 6 hanesine BIN Number diyoruz. Bu numara kart dağıtmaya yetkili banka ve kuruluşlara tahsis edilen üye numarasını temsil ediyor.

Bin numarası gereksiz gibi görünsede aslında güzel kullanıldığında ihtiyacınız olabilecek bir çok datayı barındırır. 

Örneğin ilk rakama bakarak kartın hangi servis sağlayıcısına bağlı olduğunu anlayabiliriz. 

Eğer ilk rakam
    4 ise VISA
    5 ise Mastercard
    6 ise Discover/Diners Club
    3 ise American Express/Diners Club a aittir. 

Ayrıca bu 6 rakamı kullanarak bir çok bilgi elde edebilirsiniz. Örneğin iyzico nun BIN Check servisinden dönen veriler şunlardır;

```json
{
    "status": "success",
    "locale": "tr",
    "systemTime": 1519198753973,
    "conversationId": "123456789",
    "binNumber": "542119",
    "cardType": "CREDIT_CARD",
    "cardAssociation": "MASTER_CARD",
    "cardFamily": "World",
    "bankName": "Vakıfbank",
    "bankCode": 15,
    "commercial": 0
}
```

Gelen datanın içinde kullanılan kartın tipi, servis sağlayıcısının adı, Bankanın adı ve kartın grubu gibi bilgiler bulunmakta. Ancak burada işinize yarayacak diğer bilgi bankCode ile birlikte gelmekte. 

Türk lirasını koruma kanunu gereği, Türk poslarından Türk kartları ile döviz ödemesi yapılamamaktadır, sadece Türk Lirası ile işlem yapılabilmektedir. Bundan dolayı Turkiye deki banka kartlarını bulup onları TL ye çevirip işlem yapmak isteyebilirsiniz.  Bu süreçte bankCode u kullanarak aşağıdaki bankalar ile eşleşip eşleşmediğine bakabilirsiniz. Bu şekilde kullanıcının kartının Türk kartı olup olmadığını anlayabilirsiniz.

Not : Iyzico bin servisinde sadece türkiyede faliyet gösteren bankaların bin numaraları mecvuttur.

```
046 : Akbank T.A.Ş.

143 : Aktif Yatırım Bankası A.Ş.

203 : Albaraka Türk Katılım Bankası A.Ş.

124 : Alternatifbank A.Ş.

135 : Anadolubank A.Ş.

208 : Asya Katılım Bankası A.Ş.

125 : Burgan Bank A.Ş.

092 : Citibank A.Ş.

134 : Denizbank A.Ş.

103 : Fibabanka A.Ş.

111 : Finansbank A.Ş.

062 : T. Garanti Bankası A.Ş.

012 : Halkbank A.Ş.

123 : HSBC Bank A.Ş.

109 : ICBC Turkey Bank A.Ş.

099 : ING Bank A.Ş.

205 : Kuveyt Türk Katılım Bankası A.Ş.

146 : Odea Bank A.Ş.

059 : Şekerbank T.A.Ş.

032 : Türk Ekonomi Bankası A.Ş.

096 : Turkish Bank A.Ş.

108 : Turkland Bank A.Ş.

206 : Türkiye Finans Katılım Bankası A.Ş.

064 : Türkiye İş Bankası A.Ş.

015 : T. Vakıflar Bankası T.A.O.

010 : T.C. Ziraat Bankası A.Ş.

210 : Vakıf Katılım Bankası A.Ş.

067 : Yapı ve Kredi Bankası A.Ş.

209 : Ziraat Katılım Bankası A.Ş
```