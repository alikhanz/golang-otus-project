package ucb1

import (
	"math"
)

// Получает на вход ручки и отдает индекс победившей.
func UCB1(arms Arms) int {
	var totalCount int64
	for i, arm := range arms {
		if arm.Count == 0 {
			// Отдаем предпочтение ручке которую ни разу не дергали.
			return i
		}

		totalCount += arm.Count
	}

	var res, maxRes float64
	var resultIndex int

	for i, arm := range arms {
		res = (arm.AvgIncome()) + math.Sqrt((2 * math.Log(float64(totalCount)))/float64(arm.Count))

		if res > maxRes {
			maxRes = res
			resultIndex = i
		}
	}

	return resultIndex
}