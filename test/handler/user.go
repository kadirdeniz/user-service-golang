package han_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gomock "go.uber.org/mock/gomock"
)

var testingObj *testing.T

func TestUserHandler(t *testing.T) {
	testingObj = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Handler Suite")
}

var _ = Describe("User Handler", func() {
	ctrl := gomock.NewController(testingObj)
	defer ctrl.Finish()
})
