package usecases_test

import (
	"context"
	"fmt"

	"github.com/nashmaniac/golang-application-template/adapters"
	"github.com/nashmaniac/golang-application-template/models"
	"github.com/nashmaniac/golang-application-template/usecases"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("LoginUser", func() {
	var ctx context.Context
	var u adapters.Usecases
	var username string
	var password string
	var err error
	BeforeEach(func() {
		username = "shetu"
		password = "password"
		ctx = context.Background()
		u, err = usecases.NewUsecases(ctx, mockPersistenceStore, configuration)
	})

	When("there is error in fetching from db", func() {
		It("fails", func() {
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				nil,
				fmt.Errorf("error in db"),
			).Times(1)
			_, err = u.LoginUser(ctx, username, password)
			Expect(err).ToNot(BeNil())
		})
	})

	When("the password does not match", func() {
		It("fails", func() {
			user := &models.User{
				Username: username,
				Password: "test",
			}
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				user,
				nil,
			).Times(1)
			_, err = u.LoginUser(ctx, username, password)
			Expect(err).ToNot(BeNil())
		})
	})

	When("the data is valid", func() {
		It("passes", func() {
			hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			user := &models.User{
				Username: "shetu",
				Password: string(hashedPass),
			}
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				user,
				nil,
			).Times(1)
			s, err := u.LoginUser(ctx, username, password)
			Expect(err).To(BeNil())
			Expect(s).ToNot(BeNil())
		})
	})
})
