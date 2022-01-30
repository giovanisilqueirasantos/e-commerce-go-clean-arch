package domain

import "context"

type MessageConfig struct {
	Medium            string
	From              string
	To                string
	Subject           string
	Message           string
	HasTemplate       bool
	TemplateID        string
	TemplateVariables map[string]string
}

type MessageService interface {
	SendMessage(ctx context.Context, mc *MessageConfig) error
	SendMessageFake(ctx context.Context)
}
