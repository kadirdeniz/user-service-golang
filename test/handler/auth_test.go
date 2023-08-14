package han_test

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"
	"user-service-golang/internal/dto"
	"user-service-golang/internal/entity"
	"user-service-golang/internal/handler"
	"user-service-golang/internal/service"
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

	Describe("Register", func() {
		Context("When user is not registered", func() {
			It("Should return user", func() {
				var responseObject dto.RegisterResponse

				mockRepository.EXPECT().Create(gomock.Any()).Return(mock.MockUser, nil)
				mockRepository.EXPECT().FindByEmail(gomock.Any()).Return(entity.User{}, nil)
				mockRepository.EXPECT().FindByNickname(gomock.Any()).Return(entity.User{}, nil)
				mockRepository.EXPECT().IsEmailExist(gomock.Any()).Return(false)
				mockRepository.EXPECT().IsNicknameExist(gomock.Any()).Return(false)

				handler.NewAuthHandler(
					service.NewUserService(mockRepository),
					service.NewJWT(),
					service.NewBcrypt(),
				)

				req := mock.NewMockRegisterRequest(mock.MockRegisterRequest)

				resp, err := tools.NewServer().Test(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).Should(Equal(fiber.StatusCreated))

				// read response body

				// read the response body
				body, err := io.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				fmt.Println(body)

				// unmarshal the response body into our struct
				err = json.Unmarshal(body, &responseObject)
				Expect(err).NotTo(HaveOccurred())

				fmt.Println(responseObject)

				Expect(responseObject.Token).ShouldNot(BeEmpty())
			})
		})
	})
})
