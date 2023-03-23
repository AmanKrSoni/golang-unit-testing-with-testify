package service

import (
	"fmt"

	dao_with_interface "github.com/AmanKrSoni/go-testing-example/internal/dao/with-interface"
	dao_without_interface "github.com/AmanKrSoni/go-testing-example/internal/dao/without-interface"
	with_interface "github.com/AmanKrSoni/go-testing-example/pkg/with-interface"
)

type BookServiceTestSuite struct {
	ITestMethod with_interface.ITestMethod
	IBook       dao_with_interface.IBook
}

// Not good as per Unit testing 
func NewBookServiceTestSuite() *BookServiceTestSuite {
	return &BookServiceTestSuite{IBook: dao_with_interface.NewBook()}
}


// Good as per Unit testing we can pass mock when writings test cases
func NewBookServiceTestSuitekWithDI(test with_interface.ITestMethod, bookDao dao_with_interface.IBook) *BookServiceTestSuite {
	return &BookServiceTestSuite{ITestMethod: test, IBook: bookDao}
}

// Here we have used interface method so it can be mockable
func (b *BookServiceTestSuite) GetBooksWithInterfaceForBookServiceTestSuite() {
	data := b.IBook.GetBookNamesWithInterface()
	fmt.Println("Books are : ", data)
}


// Here we called a method using pkg.Method so won`t applicable for mocking`
func (b *BookServiceTestSuite) GetBooksWithoutInterfaceForBookServiceTestSuite() {
	data := dao_without_interface.GetBookNames()
	fmt.Println("Books are : ", data)
}
