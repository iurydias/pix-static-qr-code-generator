package pix_static_qr_code_generator

import (
	"github.com/pkg/errors"
	"pix-static-qr-code-generator/payload-generator"
	"pix-static-qr-code-generator/qr-code-generator"
)

type PixStaticQRCodeGenerator struct {
	payloadGenerator payload_generator.IPayloadGenerator
	qRCodeGenerator  qr_code_generator.IQRCodeGenerator
}

func New() *PixStaticQRCodeGenerator {
	return &PixStaticQRCodeGenerator{
		payloadGenerator: payload_generator.New(),
		qRCodeGenerator:  qr_code_generator.New(),
	}
}

func (p *PixStaticQRCodeGenerator) SetData(id, pixKey, description, merchantName, merchantCity string, amount float64) {
	p.payloadGenerator.SetId(id).SetPixKey(pixKey).SetDescription(description).SetMerchantName(merchantName).SetMerchantCity(merchantCity).SetAmount(amount)
}

func (p *PixStaticQRCodeGenerator) GeneratePixQRCode(path, name string) error {
	payload, err := p.payloadGenerator.GetPayload()
	if err != nil {
		return errors.Wrap(err, "failed to generate payload")
	}
	err = p.qRCodeGenerator.Generate(payload, path+name)
	if err != nil {
		return errors.Wrap(err, "failed to generate qr code image")
	}
	return nil
}
