package withoutinterface

import (
	without_interface "github.com/AmanKrSoni/go-testing-example/pkg/without-interface"
)

func GetBookNames() []string {
	without_interface.TestWithoutInterface("TEST_WITHOUT_INTERFACE")
	return []string{"System Design", "Introducion to golang"}
}
