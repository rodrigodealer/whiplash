package handlers

import (
	"github.com/stretchr/testify/mock"
)

type esClientMock struct {
	mock.Mock
}

func (o esClientMock) Ping() int {
	args := o.Called()
	return args.Int(0)
}

func (o esClientMock) Connect() {
}

type mongoClientMock struct {
	mock.Mock
	err error
}

func (o mongoClientMock) Connect() {
}

func (o mongoClientMock) Ping() error {
	args := o.Called()
	return args.Error(0)
}
