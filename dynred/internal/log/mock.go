package log

import "fmt"

type MockRepository struct{}

func NewMockRepository() Repository {
	return &MockRepository{}
}

func (r *MockRepository) Dump(log Log) error {
	fmt.Println(log)
	return nil
}
