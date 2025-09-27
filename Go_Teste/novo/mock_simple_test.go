package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock: usar objetos falsos para teste real

type Notifier interface {
	SendNotification(message string) error
}

type MockNotifier struct {
	mock.Mock
}

func (m *MockNotifier) SendNotification(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

func NotifyUser(n Notifier, message string) error {
	return n.SendNotification(message)
}

func TestNotifyUser(t *testing.T) {
	mockNotifier := new(MockNotifier)
	mockNotifier.On("SendNotification", "Hello").Return(nil)

	err := NotifyUser(mockNotifier, "Hello")
	mockNotifier.AssertExpectations(t)
	assert.NoError(t, err)
}
