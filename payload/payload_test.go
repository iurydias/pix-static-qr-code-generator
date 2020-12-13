package payload

import (
	. "github.com/onsi/gomega"
	"strconv"
	"testing"
)

func TestNewPayload(t *testing.T) {
	RegisterTestingT(t)
	payload := NewPayload()
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
}
