package ucb1

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUCB1(t *testing.T) {
	// Заведем две ручки
	// В цикле на основе решений алгоритма UCB1 будем выбирать одну из них
	// Какая бы ручка нам не выпала, дёргаем ее и даём ей рандомно (от 0 до 1) награду
	// Первая ручка всегда будет получать награду на 0.5 больше, чем вторая.
	// В результате ожидаем, что первую ручку дернем хотя бы в 70% случаях
	// А вторую ручку хотя бы в 1% случаев
	const TRIALS = 1000

	rand.Seed(time.Now().Unix())

	arms := Arms{
		Arm{
			Count:  1,
			Reward: 1,
		},
		Arm{
			Count:  1,
			Reward: 1,
		},
	}

	hits := make(map[int]int)

	for i := 1; i <= TRIALS; i++ {
		armIndex := UCB1(arms)
		hits[armIndex]++
		arms[armIndex].Count++

		switch armIndex {
			case 0:
				arms[armIndex].Reward += rand.Float64() + 0.5
				break
			default:
				arms[armIndex].Reward += rand.Float64()
				break
		}
	}

	assert.GreaterOrEqual(t, float32(hits[0])/TRIALS*100, float32(70), "first arm must have more than 70% hits")
	assert.GreaterOrEqual(t, float32(hits[1])/TRIALS*100, float32(1), "second arm must have more than 1% hits")
}
