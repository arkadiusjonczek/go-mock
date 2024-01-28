package calculator_test

import (
	"fmt"
	"testing"

	calculator "github.com/arkadiusjonczek/go-mock"

	"go.uber.org/mock/gomock"
)

func TestAdd(t *testing.T) {
	c := calculator.NewCalculatorImpl()	
	cs := calculator.NewCalculatorService(c)
	fmt.Println(cs.Add(1))
	fmt.Println(cs.Add(1, 2))
	fmt.Println(cs.Add(1, 2, 3))
}

func TestAddWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)

	mock := NewMockCalculator(ctrl)
	mock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(float32(3)).AnyTimes()

	cs := calculator.NewCalculatorService(mock)
	
	fmt.Println(cs.Add(1)) // returns 3
	fmt.Println(cs.Add(1, 2)) // also returns 3
	fmt.Println(cs.Add(1, 2, 3)) // also returns 3
}
