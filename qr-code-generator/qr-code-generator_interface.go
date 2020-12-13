package qr_code_generator

type IQRCodeGenerator interface {
	Generate(payload, path string) error
}
