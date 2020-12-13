package payload

import (
	. "github.com/onsi/gomega"
	"strconv"
	"testing"
)

func TestNewPayload(t *testing.T) {
	RegisterTestingT(t)
	payload := New()
	t.Run("it sets an id", func(t *testing.T) {
		id := "123"
		payload.setId(id)
		Expect(payload.getId()).Should(BeEquivalentTo(id))
	})
	t.Run("it sets a pix key", func(t *testing.T) {
		key := "iury@email.com"
		payload.setPixKey(key)
		Expect(payload.getPixKey()).Should(BeEquivalentTo(key))
	})
	t.Run("it sets a description", func(t *testing.T) {
		description := "a new payment"
		payload.setDescription(description)
		Expect(payload.getDescription()).Should(BeEquivalentTo(description))
	})
	t.Run("it sets a merchant name", func(t *testing.T) {
		name := "iury"
		payload.setMerchantName(name)
		Expect(payload.getMerchantName()).Should(BeEquivalentTo(name))
	})
	t.Run("it sets a merchant city", func(t *testing.T) {
		city := "Salvador"
		payload.setMerchantCity(city)
		Expect(payload.getMerchantCity()).Should(BeEquivalentTo(city))
	})
	t.Run("it sets an amount", func(t *testing.T) {
		amount := 59.9
		payload.setAmount(amount)
		Expect(payload.getAmount()).Should(BeEquivalentTo(strconv.FormatFloat(amount, 'f', 2, 64)))
	})
	t.Run("it gets the payload", func(t *testing.T) {
		generatedPayload, err := payload.getPayload()
		Expect(err).Should(BeNil())
		Expect(generatedPayload).Should(BeEquivalentTo("00020126530014br.gov.bcb.pix0114iury@email.com0213a new payment520400005303986540559.905802BR5904iury6008Salvador62070503123630525676"))
	})
	t.Run("it gets the payload missing pix key", func(t *testing.T) {
		payload = New()
		generatedPayload, err := payload.getPayload()
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(BeEquivalentTo("pix key is missing"))
		Expect(generatedPayload).Should(BeEmpty())
	})
	t.Run("it gets the payload missing pix amount", func(t *testing.T) {
		payload = New()
		payload.setPixKey("iury@email.com")
		generatedPayload, err := payload.getPayload()
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(BeEquivalentTo("amount is missing"))
		Expect(generatedPayload).Should(BeEmpty())
	})
}
