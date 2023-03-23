package service

import "testing"

func TestGetBooksWithInterface(t *testing.T) {
	daoMock := new(MockBookDAO)
	bookSVC := NewBookMockWithDI(nil, daoMock)
	daoMock.On("GetBookNamesWithInterface").Return([]string{"TEST_MOCK"})

	bookSVC.GetBooksWithInterface()

}

// If method BookMock Service is not the has option to pass the mock pointer it will always call real method Even we have mocks
func TestGetBooksWithoutInterface(t *testing.T) {
	daoMock := new(MockBookDAO)
	bookSVC := NewBookMock()
	daoMock.On("GetBookNames").Return([]string{"TEST_MOCK"})

	bookSVC.GetBooksWithoutInterface()

}
