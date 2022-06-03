package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type messageService struct{}

func NewMessageService() *messageService {
	return &messageService{}
}

func (m messageService) SendMessage(ctx context.Context, mc *domain.MessageConfig) error {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration((8 + rand.Intn(5))) * time.Second)

	return nil
}

func (m messageService) SendMessageFake(ctx context.Context) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration((8 + rand.Intn(5))) * time.Second)
}
