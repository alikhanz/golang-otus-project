package ucb1

type Arm struct {
	Count int64  // количество "дерганий за ручку"
	Reward float64 // положительное подкрепление
}

func (a *Arm) AvgIncome() float64 {
	return a.Reward / float64(a.Count)
}

type Arms []Arm