package withdrawal

//DefaultRepository sample basic repository for withdrawals
type DefaultRepository struct {
}

//GetAvailableBanknotesValues returns available notes on ATM
func (d DefaultRepository) GetAvailableBanknotesValues() []int {
	return []int{50, 10, 5, 1}
}

//CreateDefaultRepository returns a new instance of DefaultRepository
func CreateDefaultRepository() BaseRepository {
	return DefaultRepository{}
}
