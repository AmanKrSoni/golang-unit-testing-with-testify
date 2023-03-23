package withinterface

import "fmt"

type ITestMethod interface {
	TestWithInterface(val string)
}

type TestMethodImpl struct{}

func NewTestMethod() ITestMethod {
	return &TestMethodImpl{}
}

func (t *TestMethodImpl) TestWithInterface(val string) {
	fmt.Println("value passed to TestWithInterface() -> ", val)
}
