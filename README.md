# PIX Static QR Code Generator

A PIX static QR Code generator to help in golang services development 

## Usage

#### Setting data to PIX Qr Code

##### Example

```go
    pix_static_qr_code_generator.SetData("123", "iury@email.com", "a new payment", "Iury", "Salvador", 15.00)
```
 ##### Setting Data
 
 ```go
    // id is a identification for the transaction *optional
    // pixKey is the receiver PIX key *required
    // description is a additional information for the transaction *optional
    // merchantName is the name of the PIX receiver *optional
    // merchantCity is the city of the PIX receiver *optional
    // amount is the amount of the transaction *required
     pix_static_qr_code_generator.SetData(id, pixKey, description, merchantName, merchantCity string, amount float64)
 ```
 
#### Generating QR Code

##### Example
```go
   pix_static_qr_code_generator.GeneratePixQRCode("./", "qr")
```
##### Generating
 ```go
    // path is where the qr code image will be saved, with a slash in the end
    // name is the qr code file name
     pix_static_qr_code_generator.GeneratePixQRCode(path, name string)
 ```