package service_test

import (
	"testing"
	"user-service-golang/internal/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testingObj *testing.T

func TestBcryptService(t *testing.T) {
	testingObj = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bcrypt Suite")
}

var _ = Describe("Bcrypt Service", func() {
	bcryptService := service.NewBcrypt()

	It("Should return a hashed password", func() {
		hashedPassword := bcryptService.HashPassword("password")
		Expect(hashedPassword).ToNot(BeNil())
	})

	It("Should check is passwords match", func() {
		hashedPassword := bcryptService.HashPassword("password")
		isMatch := bcryptService.IsPasswordCorrect(hashedPassword, "password")
		Expect(isMatch).To(BeTrue())
	})
})
