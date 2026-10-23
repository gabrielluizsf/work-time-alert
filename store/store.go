package store

type Store interface {
	Save(alerts ...Alert) error
	Load() ([]Alert, error)
}