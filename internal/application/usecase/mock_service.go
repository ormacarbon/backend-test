package usecase

import (
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/eventbus"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/smtp"
	"github.com/stretchr/testify/mock"
)

type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) Send(input smtp.InputSendEmail) error {
	args := m.Called(input)
	return args.Error(0)
}

type MockEventBus struct {
	mock.Mock
}

func NewMockEventBus() *MockEventBus {
	return &MockEventBus{}
}

func (m *MockEventBus) Subscribe(eventName string, subscriber chan<- eventbus.Event) error {
	args := m.Called(eventName, subscriber)
	return args.Error(0)
}

func (m *MockEventBus) Publish(event eventbus.Event) {
	m.Called(event)
}
