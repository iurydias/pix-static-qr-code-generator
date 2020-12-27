package payload_generator

import (
	"errors"
	"fmt"
	"github.com/snksoft/crc"
	"strconv"
)

type PayloadGenerator struct {
	id          string
	pixKey      string
	description string
	merchant    merchant
	amount      string
}

type merchant struct {
	name string
	city string
}

func New() *PayloadGenerator {
	return new(PayloadGenerator)
}

func (p *PayloadGenerator) GetPayload() (string, error) {
	if p.pixKey == "" {
		return "", errors.New("pix key is missing")
	}
	payloadFormatIndicator := p.getValue(PAYLOAD_FORMAT_INDICATOR, "01")
	merchantAccountInformation := p.getMerchantAccountInformation()
	merchantCategoryCode := p.getValue(MERCHANT_CATEGORY_CODE, "0000")
	transactionCode := p.getValue(TRANSACTION_CURRENCY, "986")
	transactionAmount := p.getValue(TRANSACTION_AMOUNT, p.amount)
	merchantName := p.getValue(MERCHANT_NAME, p.merchant.name)
	merchantCity := p.getValue(MERCHANT_CITY, p.merchant.city)
	countryCode := p.getValue(COUNTRY_CODE, "BR")
	additionalData := p.getAdditionalDataFieldTemplate()
	payload := payloadFormatIndicator +
		merchantAccountInformation +
		merchantCategoryCode +
		transactionCode +
		transactionAmount +
		countryCode +
		merchantName +
		merchantCity +
		additionalData
	crcValue := p.getValue(CRC16, p.getCRC(payload))
	return payload + crcValue, nil
}

func (p *PayloadGenerator) getCRC(payload string) string {
	return strconv.Itoa(int(crc.CalculateCRC(crc.CCITT, []byte(payload))))
}

func (p *PayloadGenerator) getAdditionalDataFieldTemplate() string {
	txId := p.getValue(ADDITIONAL_DATA_FIELD_TEMPLATE_TXID, p.id)
	return p.getValue(ADDITIONAL_DATA_FIELD_TEMPLATE, txId)
}

func (p *PayloadGenerator) getMerchantAccountInformation() string {
	merchantGUI := p.getValue(MERCHANT_ACCOUNT_INFORMATION_GUI, "br.gov.bcb.pix")
	merchantKey := p.getValue(MERCHANT_ACCOUNT_INFORMATION_KEY, p.pixKey)
	merchantDescription := p.getValue(MERCHANT_ACCOUNT_INFORMATION_DESCRIPTION, p.description)
	merchantAccountInformation := merchantGUI + merchantKey + merchantDescription
	return p.getValue(MERCHANT_ACCOUNT_INFORMATION, merchantAccountInformation)
}

func (p *PayloadGenerator) getValue(id, content string) string {
	return fmt.Sprintf("%s%02d%s", id, len(content), content)
}
