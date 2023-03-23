package main

import (
	"fmt"

	dao_with_interface "github.com/AmanKrSoni/go-testing-example/internal/dao/with-interface"
	svc "github.com/AmanKrSoni/go-testing-example/internal/service"
	with_interface "github.com/AmanKrSoni/go-testing-example/pkg/with-interface"
)

func main() {
	fmt.Println("starting main .... ")

	// created an instace of bookMock without passing any reference
	bookMock := svc.NewBookMock()
	bookMock.GetBooksWithoutInterface()

	bookMockWithDI := svc.NewBookMockWithDI(with_interface.NewTestMethod(), dao_with_interface.NewBook())
	bookMockWithDI.GetBooksWithInterface()

}
