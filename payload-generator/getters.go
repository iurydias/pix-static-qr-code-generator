package payload_generator

func (p *PayloadGenerator) GetId() string {
	return p.id
}

func (p *PayloadGenerator) GetPixKey() string {
	return p.pixKey
}

func (p *PayloadGenerator) GetDescription() string {
	return p.description
}

func (p *PayloadGenerator) GetMerchantName() string {
	return p.merchant.name
}

func (p *PayloadGenerator) GetMerchantCity() string {
	return p.merchant.city
}

func (p *PayloadGenerator) GetAmount() string {
	return p.amount
}
