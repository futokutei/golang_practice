package calc

import "errors"

func Sum(nums ...float64) float64 {
	res := 0.0

	for _, num := range nums {
		res += num
	}
	return res
}

func Max(nums ...float64) float64 {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}

func Min(nums ...float64) float64 {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Can`t division by 0")
	}
	return a / b, nil
}
