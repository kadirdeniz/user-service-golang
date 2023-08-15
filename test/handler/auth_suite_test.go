package han_test

import (
	"fmt"
	"testing"
	"user-service-golang/internal/dto"
	"user-service-golang/internal/handler"
	"user-service-golang/internal/service"
	"user-service-golang/pkg"
	"user-service-golang/test"
	"user-service-golang/test/mock"
	"user-service-golang/tools"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gomock "go.uber.org/mock/gomock"
)

var testingObj *testing.T

func TestAuthHandler(t *testing.T) {
	testingObj = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Handler Suite")
}

var _ = Describe("Auth Handler", func() {
	ctrl := gomock.NewController(testingObj)
	defer ctrl.Finish()

	var mockRepository = mock.NewMockUserRepositoryActions(ctrl)
	Describe("Register", Ordered, func() {
		Context("When user is not registered", func() {
			It("Should register user", func() {
				mockRepository.EXPECT().Create(gomock.Any()).Return(mock.MockUser, nil)
				mockRepository.EXPECT().IsEmailExist(gomock.Any()).Return(false)
				mockRepository.EXPECT().IsNicknameExist(gomock.Any()).Return(false)

				handler.NewAuthHandler(
					service.NewAuthService(
						service.NewUserService(mockRepository),
					),
				)

				req := mock.NewMockRegisterRequest(mock.MockRegisterRequest)

				resp, err := tools.NewServer().Test(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).Should(Equal(fiber.StatusCreated))

				var responseObject dto.TokenResponse

				test.ReadResponseBody(resp.Body, &responseObject)

				Expect(responseObject.Token).ShouldNot(BeEmpty())
			})
		})

		Context("When user is registered", func() {
			When("Email is exist", func() {
				It("Should return error", func() {
					mockRepository.EXPECT().IsEmailExist(gomock.Any()).AnyTimes().Return(true)
					mockRepository.EXPECT().IsNicknameExist(gomock.Any()).AnyTimes().Return(false)

					handler.NewAuthHandler(
						service.NewAuthService(
							service.NewUserService(mockRepository),
						),
					)

					req := mock.NewMockRegisterRequest(mock.MockRegisterRequest)

					resp, err := tools.NewServer().Test(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).Should(Equal(fiber.StatusBadRequest))

					var responseObject pkg.ErrorResponse

					test.ReadResponseBody(resp.Body, &responseObject)

					Expect(responseObject.Message).Should(Equal(pkg.ErrEmailAlreadyExist.Error()))
				})
			})

			When("Nickname is exist", func() {
				It("Should return error", func() {
					mockRepository.EXPECT().IsEmailExist(gomock.Any()).AnyTimes().Return(false)
					mockRepository.EXPECT().IsNicknameExist(gomock.Any()).AnyTimes().Return(true)

					handler.NewAuthHandler(
						service.NewAuthService(
							service.NewUserService(mockRepository),
						),
					)

					req := mock.NewMockRegisterRequest(mock.MockRegisterRequest)

					resp, err := tools.NewServer().Test(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).Should(Equal(fiber.StatusBadRequest))

					var responseObject pkg.ErrorResponse

					test.ReadResponseBody(resp.Body, &responseObject)

					Expect(responseObject.Message).Should(Equal(pkg.ErrEmailAlreadyExist.Error()))
				})
			})
		})
	})

	Describe("Login", func() {
		Context("When user is registered", func() {
			It("Should login user", func() {
				mockRepository.EXPECT().FindByEmail(gomock.Any()).Return(mock.MockUser, nil)

				handler.NewAuthHandler(
					service.NewAuthService(
						service.NewUserService(mockRepository),
					),
				)

				req := mock.NewMockLoginRequest(mock.MockLoginRequest)

				resp, err := tools.NewServer().Test(req)
				fmt.Println(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).Should(Equal(fiber.StatusOK))

				var responseObject dto.TokenResponse

				test.ReadResponseBody(resp.Body, &responseObject)

				Expect(responseObject.Token).ShouldNot(BeEmpty())
			})
		})

		Context("When user is not registered", func() {
			It("Should return error", func() {
				mockRepository.EXPECT().FindByEmail(gomock.Any()).Return(mock.MockUser, pkg.ErrUserNotFound)

				handler.NewAuthHandler(
					service.NewAuthService(
						service.NewUserService(mockRepository),
					),
				)

				req := mock.NewMockLoginRequest(mock.MockLoginRequest)

				resp, err := tools.NewServer().Test(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).Should(Equal(fiber.StatusBadRequest))

				var responseObject pkg.ErrorResponse

				test.ReadResponseBody(resp.Body, &responseObject)

				Expect(responseObject.Message).Should(Equal(pkg.ErrUserNotFound.Error()))
			})
		})
	})
})
