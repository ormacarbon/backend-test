package internal

import "github.com/stretchr/testify/mock"

type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) Send(input InputSendEmail) error {
	args := m.Called(input)
	return args.Error(0)
}

type MockEventBus struct {
	mock.Mock
}

func NewMockEventBus() *MockEventBus {
	return &MockEventBus{}
}

func (m *MockEventBus) Subscribe(eventName string, subscriber chan<- Event) error {
	args := m.Called(eventName, subscriber)
	return args.Error(0)
}

func (m *MockEventBus) Publish(event Event) {
	m.Called(event)
}
