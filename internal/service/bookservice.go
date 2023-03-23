package service

import (
	"fmt"

	dao_with_interface "github.com/AmanKrSoni/go-testing-example/internal/dao/with-interface"
	dao_without_interface "github.com/AmanKrSoni/go-testing-example/internal/dao/without-interface"
	with_interface "github.com/AmanKrSoni/go-testing-example/pkg/with-interface"
)

type BookService struct {
	ITestMethod with_interface.ITestMethod
	IBook       dao_with_interface.IBook
}

func NewBookMock() *BookService {
	return &BookService{IBook: dao_with_interface.NewBook()}
}

func NewBookMockWithDI(test with_interface.ITestMethod, bookDao dao_with_interface.IBook) *BookService {
	return &BookService{ITestMethod: test, IBook: bookDao}
}

func (b *BookService) GetBooksWithInterface() {
	data := b.IBook.GetBookNamesWithInterface()
	fmt.Println("Books are : ", data)
}

func (b *BookService) GetBooksWithoutInterface() {
	data := dao_without_interface.GetBookNames()
	fmt.Println("Books are : ", data)
}
