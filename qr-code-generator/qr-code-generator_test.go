package qr_code_generator

import (
	"bytes"
	"github.com/liyue201/goqr"
	. "github.com/onsi/gomega"
	"image"
	"io/ioutil"
	"testing"
)

func TestQRCodeGenerator(t *testing.T) {
	RegisterTestingT(t)
	generator := New()
	t.Run("it generates a qr code image", func(t *testing.T) {
		payload := "payload-generator-test"
		path := "./qr.png"
		err := generator.Generate(payload, path)
		Expect(err).ShouldNot(HaveOccurred())
		img := getImage(path)
		qrCode, err := goqr.Recognize(img)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(len(qrCode)).Should(BeEquivalentTo(1))
		Expect(qrCode[0].Payload).Should(BeEquivalentTo(payload))
	})
}

func getImage(path string) image.Image {
	imgData, err := ioutil.ReadFile(path)
	Expect(err).ShouldNot(HaveOccurred())
	img, _, err := image.Decode(bytes.NewReader(imgData))
	Expect(err).ShouldNot(HaveOccurred())
	return img
}
