package qr_code_generator

import "github.com/skip2/go-qrcode"

type QRCodeGenerator struct{}

func New() *QRCodeGenerator {
	return new(QRCodeGenerator)
}

func (q *QRCodeGenerator) Generate(payload, path string) error {
	return qrcode.WriteFile(payload, qrcode.Medium, 256, path)
}
