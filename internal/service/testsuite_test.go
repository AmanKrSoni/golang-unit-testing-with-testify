package service

import (
	"fmt"
	"testing"

	dao_with_interface "github.com/AmanKrSoni/go-testing-example/internal/dao/with-interface"
	mock_dao_with_interface "github.com/AmanKrSoni/go-testing-example/internal/dao/with-interface/mocks"
	_ "github.com/AmanKrSoni/go-testing-example/pkg/with-interface"
	mock_with_interface "github.com/AmanKrSoni/go-testing-example/pkg/with-interface/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestSuiteForBookServiceTestSuiteStruct struct {
	suite.Suite
	service       *BookServiceTestSuite
	service2      *BookServiceTestSuite
	ITestMethod   *mock_with_interface.ITestMethod
	BookInterface dao_with_interface.IBook
	IBook         *mock_dao_with_interface.IBook
}

func (suite *TestSuiteForBookServiceTestSuiteStruct) SetupTest() {
	fmt.Printf("\n\n --->>> !!! Setup is called for suite : TestSuiteForBookServiceTestSuiteStruct !!!\n\n")

	suite.ITestMethod = new(mock_with_interface.ITestMethod)
	suite.IBook = new(mock_dao_with_interface.IBook)
	suite.service = NewBookServiceTestSuitekWithDI(suite.ITestMethod, suite.IBook)
	suite.BookInterface = dao_with_interface.NewBookWithDI(suite.ITestMethod)
	suite.service2 = NewBookServiceTestSuitekWithDI(suite.ITestMethod, suite.IBook)
}

// This test is perform for outer method mocking i.e IBook Methods
func (suite *TestSuiteForBookServiceTestSuiteStruct) TestGetBooksWithInterfaceForBookServiceTestSuite() {
	suite.ITestMethod.On("TestWithInterface", mock.Anything)
	suite.IBook.On("GetBookNamesWithInterface").Return([]string{"TEST_SUITE_MOCK_RETURNS"})
	suite.service.GetBooksWithInterfaceForBookServiceTestSuite()

}

// This test is perform for Inner method mocking i.e ITestMethod Methods mocksing
func (suite *TestSuiteForBookServiceTestSuiteStruct) TestGetBooksWithInterfaceForBookServiceTestSuiteInnerMethodMock() {
	suite.ITestMethod.On("TestWithInterface", mock.Anything)
	suite.IBook.On("GetBookNamesWithInterface").Return([]string{"TEST_SUITE_MOCK_RETURNS"})
	suite.service2.GetBooksWithInterfaceForBookServiceTestSuite()

}

func (suite *TestSuiteForBookServiceTestSuiteStruct) TearDownTest() {
	fmt.Printf("\n\n--->>> !!! Tear Down the test suite : TestSuiteForBookServiceTestSuiteStruct !!!\n\n\n\n\n")
}

func TestRun(t *testing.T) {
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("\n\n Running TestSuiteForBookServiceTestSuiteStruct !!! \n")
	suite.Run(t, new(TestSuiteForBookServiceTestSuiteStruct))
}
