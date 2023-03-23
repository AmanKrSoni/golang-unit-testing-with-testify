package service

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

// Created a Mock Obj which can be subsitutable to actual implementations
type MockBookDAO struct {
	mock.Mock
}

// Implementing interface methods
func (m *MockBookDAO) GetBookNamesWithInterface() []string {
	args := m.Called()
	fmt.Println("Mock Dao GetBookNamesWithInterface() is called ....")
	//return []string{"MOCK"}
	var r0 []string
	if rf, ok := args.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]string)
		}
	}
	return r0
}

func (m *MockBookDAO) GetAuthor() []string {
	ret := m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}
