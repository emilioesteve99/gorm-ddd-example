package commonHttpModelFixtures

import (
	"context"
	testUtils "gorm-ddd-example/src/common/fixture/utils"
)

type HttpRequestMock struct {
	context.Context
	testUtils.BaseMock
}

func (m *HttpRequestMock) JSON(code int, obj any) {
	m.Called(code, obj)
}
