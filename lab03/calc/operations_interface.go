package calc

import "errors"

type Calcutator interface {
	Sum(nums ...float64) float64
	Max(nums ...float64) float64
	Min(nums ...float64) float64
	Divide(a, b float64) (float64, error)
}

type Calc struct{}

func (c Calc) Sum(nums ...float64) float64 {
	res := 0.0

	for _, num := range nums {
		res += num
	}
	return res
}

func (c Calc) Max(nums ...float64) float64 {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}

func (c Calc) Min(nums ...float64) float64 {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}

func (c Calc) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Can`t division by 0")
	}
	return a / b, nil
}
