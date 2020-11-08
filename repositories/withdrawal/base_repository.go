package withdrawal

//BaseRepository defines basic methods
type BaseRepository interface {
	GetAvailableBanknotesValues() []int16
}
