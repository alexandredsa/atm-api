package withdrawal

type DefaultRepository struct {
}

func (d DefaultRepository) GetAvailableBanknotesValues() []int {
	return []int{50, 10, 5, 1}
}

func CreateDefaultRepository() BaseRepository {
	return DefaultRepository{}
}
