package handlers

import (
	"github.com/stretchr/testify/mock"
)

type clientMock struct {
	mock.Mock
}

func (o clientMock) Ping() int {
	args := o.Called()
	return args.Int(0)
}

func (o clientMock) Connect() {
}
