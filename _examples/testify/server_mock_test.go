/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package testify

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"
)

// serverMock mock
type serverMock struct {
	mock.Mock
}

func NewServerMock() *serverMock {
	return &serverMock{}
}

// Serve mocked method
func (m *serverMock) Serve(p0 string) ([]byte, error) {

	ret := m.Called(p0)

	var r0 []byte
	switch res := ret.Get(0).(type) {
	case nil:
	case []byte:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}
