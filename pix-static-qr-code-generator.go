package pix_static_qr_code_generator

import (
	"github.com/iiurydias/pix-static-qr-code-generator/payload-generator"
	"github.com/iiurydias/pix-static-qr-code-generator/qr-code-generator"
	"github.com/pkg/errors"
)

type PixStaticQRCodeGenerator struct {
	payloadGenerator *payload_generator.PayloadGenerator
	qRCodeGenerator  *qr_code_generator.QRCodeGenerator
}

func New(pixKey string) *PixStaticQRCodeGenerator {
	return &PixStaticQRCodeGenerator{
		payloadGenerator: payload_generator.New().SetPixKey(pixKey),
		qRCodeGenerator:  qr_code_generator.New(),
	}
}

func (p *PixStaticQRCodeGenerator) SetId(id string) {
	p.payloadGenerator.SetId(id)
}

func (p *PixStaticQRCodeGenerator) SetAmount(amount float64) {
	p.payloadGenerator.SetAmount(amount)
}

func (p *PixStaticQRCodeGenerator) SetDescription(description string) {
	p.payloadGenerator.SetDescription(description)
}

func (p *PixStaticQRCodeGenerator) SetMerchantName(merchantName string) {
	p.payloadGenerator.SetMerchantName(merchantName)
}

func (p *PixStaticQRCodeGenerator) SetMerchantCity(merchantCity string) {
	p.payloadGenerator.SetMerchantCity(merchantCity)
}

func (p *PixStaticQRCodeGenerator) Generate(path string) error {
	payload, err := p.payloadGenerator.GetPayload()
	if err != nil {
		return errors.Wrap(err, "failed to generate payload")
	}
	if err = p.qRCodeGenerator.Generate(payload, path); err != nil {
		return errors.Wrap(err, "failed to generate qr code image")
	}
	return nil
}
