package payload_generator

import "strconv"

func (p *PayloadGenerator) SetId(id string) *PayloadGenerator {
	p.id = id
	return p
}

func (p *PayloadGenerator) SetPixKey(key string) *PayloadGenerator {
	p.pixKey = key
	return p
}

func (p *PayloadGenerator) SetDescription(description string) *PayloadGenerator {
	p.description = description
	return p
}

func (p *PayloadGenerator) SetMerchantName(name string) *PayloadGenerator {
	p.merchant.name = name
	return p
}

func (p *PayloadGenerator) SetMerchantCity(city string) *PayloadGenerator {
	p.merchant.city = city
	return p
}

func (p *PayloadGenerator) SetAmount(amount float64) *PayloadGenerator {
	p.amount = strconv.FormatFloat(amount, 'f', 2, 64)
	return p
}
