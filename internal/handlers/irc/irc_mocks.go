package irc

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type MockedClient struct {
	mock.Mock
	Client
}

func (mc *MockedClient) addHandlers() {}

func (mc *MockedClient) Connect() error {
	return errors.New("some error")
}
