package payload_generator

type IPayloadGenerator interface {
	GetPayload() (string, error)
	SetId(id string) *PayloadGenerator
	SetPixKey(key string) *PayloadGenerator
	SetDescription(description string) *PayloadGenerator
	SetMerchantName(name string) *PayloadGenerator
	SetMerchantCity(city string) *PayloadGenerator
	SetAmount(amount float64) *PayloadGenerator
	GetId() string
	GetPixKey() string
	GetDescription() string
	GetMerchantName() string
	GetMerchantCity() string
	GetAmount() string
}
