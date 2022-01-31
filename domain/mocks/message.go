package mocks

import (
	"context"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockMessageService struct {
	mock.Mock
}

func (mms *MockMessageService) SendMessage(ctx context.Context, mc *domain.MessageConfig) error {
	args := mms.Called(ctx, mc)
	return args.Error(0)
}

func (mms *MockMessageService) SendMessageFake(ctx context.Context) {}
