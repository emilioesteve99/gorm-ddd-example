package testUtils

import "github.com/stretchr/testify/mock"

type BaseMock struct {
	mock.Mock
}

func (m *BaseMock) ResetMock() {
	m.ExpectedCalls = nil
	m.Calls = nil
}
