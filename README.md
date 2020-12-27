# PIX Static QR Code Generator

A PIX static QR Code generator to help in golang services development 

## Usage

### Instancing a new generator

```go
    // passing the pix key
    generator := pix_static_qr_code_generator.New("iury@email.com")
```

### Setting data to PIX Qr Code

#### Methods available

***.SetAmount()***

It sets an amount to the payment

**Example**
```go
    generator.SetAmount(15.00)
```
***.SetDescription()***

It sets a description to the payment

**Example**
```go
    generator.SetDescription("a new payment")
```

***.SetId()***

It sets an identification to the payment

**Example**
```go
    generator.SetId("123")
```
***.SetMerchantName()***

It sets a merchant name to the payment

**Example**
```go
    generator.SetMerchantName("Iury")
```
***.SetMerchantCity()***

It sets a merchant city to the payment

**Example**
```go
    generator.SetMerchantCity("Salvador")
```
#### Generating QR Code

##### Example

```go
    // passing the path where the qr code image will be saved and your file name
   generator.Generate("./qr.png")
```
