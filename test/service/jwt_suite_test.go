package service_test

import (
	"testing"
	"user-service-golang/internal/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJWTHandler(t *testing.T) {
	testingObj = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "JWT Suite")
}

var _ = Describe("JWT Service", func() {
	var jwt service.JWTActions

	BeforeEach(func() {
		jwt = service.NewJWT()
	})

	Context("Set User ID", func() {
		It("should set user id", func() {
			jwt.SetUserId(uint(1))
			Expect(jwt.GetUserId()).To(Equal(uint(1)))
		})
	})

	Context("Set Token", func() {
		It("should set token", func() {
			jwt.SetToken("token")
			Expect(jwt.GetToken()).To(Equal("token"))
		})
	})

	Context("Generate Token", func() {
		It("should parse token", func() {
			mockToken := jwt.SetUserId(9).CreateToken()
			claims := mockToken.ParseToken()
			Expect(claims.GetUserId()).To(Equal(uint(9)))
		})
	})
})
