package payload

import "strconv"

func (p *Payload) setId(id string) *Payload {
	p.id = id
	return p
}

func (p *Payload) setPixKey(key string) *Payload {
	p.pixKey = key
	return p
}

func (p *Payload) setDescription(description string) *Payload {
	p.description = description
	return p
}

func (p *Payload) setMerchantName(name string) *Payload {
	p.merchant.name = name
	return p
}

func (p *Payload) setMerchantCity(city string) *Payload {
	p.merchant.city = city
	return p
}

func (p *Payload) setAmount(amount float64) *Payload {
	p.amount = strconv.FormatFloat(amount, 'f', 2, 64)
	return p
}
