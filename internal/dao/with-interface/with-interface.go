package withinterface

import (
	with_interface "github.com/AmanKrSoni/go-testing-example/pkg/with-interface"
)

type IBook interface {
	GetBookNamesWithInterface() []string
	GetAuthor() []string
}

type BookImpl struct {
	Test with_interface.ITestMethod
}

func NewBook() IBook {
	return &BookImpl{Test: with_interface.NewTestMethod()}
}

func NewBookWithDI(test with_interface.ITestMethod) IBook {
	return &BookImpl{Test: test}
}

func (b *BookImpl) GetBookNamesWithInterface() []string {
	b.Test.TestWithInterface("TEST_WITH_INTERFACE")
	return []string{"System Design", "Introducion to golang"}
}

func (b *BookImpl) GetAuthor() []string {
	return []string{"Author1", "Author2"}
}
