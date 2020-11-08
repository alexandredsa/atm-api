package withdrawal

type DefaultRepository struct {
}

func (d DefaultRepository) GetAvailableBanknotesValues() []int16 {
	return []int16{50, 10, 5, 1}
}
