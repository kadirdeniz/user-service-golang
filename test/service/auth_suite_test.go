package service_test

import (
	"testing"
	"user-service-golang/internal/service"
	"user-service-golang/test/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

func TestAuthService(t *testing.T) {
	testingObj = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Service")
}

var _ = Describe("Auth Service", func() {
	ctrl := gomock.NewController(testingObj)
	defer ctrl.Finish()

	var mockRepository = mock.NewMockUserRepositoryActions(ctrl)
	var authService service.AuthServiceActions

	BeforeEach(func() {
		authService = service.NewAuthService(
			service.NewUserService(mockRepository),
		)
	})

	Context("Register", func() {
		When("user is valid", func() {
			It("should return token", func() {
				mockRepository.EXPECT().Create(gomock.Any()).Return(mock.MockUser, nil).Times(1)
				mockRepository.EXPECT().IsEmailExist(gomock.Any()).Return(false).Times(1)
				mockRepository.EXPECT().IsNicknameExist(gomock.Any()).Return(false).Times(1)

				token, err := authService.Register(mock.MockUser)
				Expect(err).To(BeNil())
				Expect(token).ToNot(BeEmpty())
			})
		})

		When("user is invalid", func() {
			It("should return error", func() {
				mockRepository.EXPECT().Create(gomock.Any()).Return(mock.MockUser, nil).Times(1)
				mockRepository.EXPECT().IsEmailExist(gomock.Any()).Return(true).Times(1)

				token, err := authService.Register(mock.MockUser)
				Expect(err).ToNot(BeNil())
				Expect(token).To(BeEmpty())
			})
		})
	})

	Context("Login", func() {
		When("user is valid", func() {
			It("should return token", func() {
				mockRepository.EXPECT().FindByEmail(gomock.Any()).Return(mock.MockUser, nil).Times(1)

				token, err := authService.Login(mock.MockUser)
				Expect(err).To(BeNil())
				Expect(token).ToNot(BeEmpty())
			})
		})

		When("user is invalid", func() {
			It("should return error", func() {
				mockRepository.EXPECT().FindByEmail(gomock.Any()).Return(mock.MockUser, nil).Times(1)
				mockRepository.EXPECT().IsEmailExist(gomock.Any()).Return(false).Times(1)

				token, err := authService.Login(mock.MockUser)
				Expect(err).ToNot(BeNil())
				Expect(token).To(BeEmpty())
			})
		})
	})
})
