package usecases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nashmaniac/golang-application-template/config"
	"github.com/nashmaniac/golang-application-template/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var mockPersistenceStore *mocks.MockPeristenceStore
var mockController *gomock.Controller
var configuration *config.Config

func TestUsecases(t *testing.T) {
	RegisterFailHandler(Fail)

	configuration = &config.Config{
		SecretKey: "hello-world",
	}
	mockController = gomock.NewController(GinkgoT())
	mockPersistenceStore = mocks.NewMockPeristenceStore(mockController)

	RunSpecs(t, "Usecases Suite")
}
