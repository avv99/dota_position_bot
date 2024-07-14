package storage

type Storage interface {
	GetHeroes(string) ([]string, error)
}
