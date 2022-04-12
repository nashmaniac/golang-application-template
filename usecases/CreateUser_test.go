package usecases_test

import (
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/nashmaniac/golang-application-template/adapters"
	"github.com/nashmaniac/golang-application-template/models"
	"github.com/nashmaniac/golang-application-template/usecases"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create User", func() {
	var ctx context.Context
	var u adapters.Usecases
	var username string
	var password string
	var err error
	BeforeEach(func() {
		username = "shetu"
		password = "password"
		ctx = context.Background()
		u, err = usecases.NewUsecases(ctx, mockPersistenceStore)
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
			_, err = u.CreateUser(ctx, username, password)
			Expect(err).ToNot(BeNil())
		})
	})

	When("there is no user found", func() {
		It("fails", func() {
			user := &models.User{
				Username: username,
				Password: password,
			}
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				user,
				nil,
			).Times(1)
			_, err = u.CreateUser(ctx, username, password)
			Expect(err).ToNot(BeNil())
		})
	})
	When("there is error in creating user", func() {
		It("fails", func() {
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				nil,
				nil,
			).Times(1)

			mockPersistenceStore.EXPECT().CreateUser(
				ctx,
				gomock.Any(),
			).Return(
				nil,
				fmt.Errorf("error in db"),
			).Times(1)
			_, err = u.CreateUser(ctx, username, password)
			Expect(err).ToNot(BeNil())
		})
	})
	When("there is no error in creating user", func() {
		It("passes", func() {
			user := &models.User{
				Username: username,
				Password: password,
			}
			mockPersistenceStore.EXPECT().FindUserByUsername(
				ctx,
				username,
			).Return(
				nil,
				nil,
			).Times(1)

			mockPersistenceStore.EXPECT().CreateUser(
				ctx,
				gomock.Any(),
			).Return(
				user,
				nil,
			).Times(1)
			user, err = u.CreateUser(ctx, username, password)
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
		})
	})
})
