package payload

func (p *Payload) getId() string {
	return p.id
}

func (p *Payload) getPixKey() string {
	return p.pixKey
}

func (p *Payload) getDescription() string {
	return p.description
}

func (p *Payload) getMerchantName() string {
	return p.merchant.name
}

func (p *Payload) getMerchantCity() string {
	return p.merchant.city
}

func (p *Payload) getAmount() string {
	return p.amount
}
