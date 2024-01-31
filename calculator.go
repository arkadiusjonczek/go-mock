package calculator

//go:generate mockgen -source=calculator.go -destination=mock_calculator_test.go -package=calculator_test

import (
	"fmt"
)

type Calculator interface {
	Add(a, b float32) float32
}

// CalculatorImpl implements Calculator
var _ Calculator = (*CalculatorImpl)(nil)

type CalculatorImpl struct {
}

func NewCalculatorImpl() Calculator {
	return &CalculatorImpl{}
}

func (c *CalculatorImpl) Add(a, b float32) float32 {
	return a + b
}

type CalculatorService struct {
	calculator Calculator
}

func NewCalculatorService(calculator Calculator) *CalculatorService {
	return &CalculatorService{
		calculator: calculator,
	}
}

func (cs *CalculatorService) Add(numbers ...float32) (float32, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers given for addition")
	}
	
	result := float32(0)

	for _, number := range numbers {
		result = cs.calculator.Add(result, number)
	}
	
	return result, nil
}


