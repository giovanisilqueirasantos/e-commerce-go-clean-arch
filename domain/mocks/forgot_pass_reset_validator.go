package mocks

import (
	"context"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockForgotPassResetValidator struct {
	mock.Mock
}

func (mfprv *MockForgotPassResetValidator) Validate(ctx context.Context, fpr *domain.ForgotPassReset) (domain.IsValid, domain.Message, error) {
	args := mfprv.Called(ctx, fpr)
	return domain.IsValid(args.Bool(0)), domain.Message(args.String(1)), args.Error(2)
}

