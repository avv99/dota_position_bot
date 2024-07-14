package inmemory

import (
	"dota_position_bot/internal/storage"
	"errors"
	"log"
)

type Inmemory struct {
	Heroes map[string][]string
}

func InitInMemoryStorage() (storage.Storage, error) {
	var List Inmemory
	List.LoadData()
	return &List, nil
}

func (i *Inmemory) GetHeroes(pos string) ([]string, error) {
	heroes, ok := i.Heroes[pos]
	if ok == false {
		log.Println("Выбрана неправильная позиция")
		return []string{}, errors.New("Выбрана неправильная позиция")
	}
	return heroes, nil
}

func (i *Inmemory) LoadData() {
	i.Heroes = make(map[string][]string)
	i.Heroes["1pos"] = []string{"Windranger", "Broodmother", "Clinkz", "MonkeyKing"}
	i.Heroes["2pos"] = []string{"Tiny", "Broodmother", "EmberSpirit", "PrimalBeast"}
	i.Heroes["3pos"] = []string{"DarkSeer", "Broodmother", "Enigma", "Windranger"}
	i.Heroes["4pos"] = []string{"SpiritBreaker", "Clockwerk", "ShadowDemon", "Weawer"}
	i.Heroes["5pos"] = []string{"Sven", "ElderTitan", "TreantProtector", "Weawer"}
}
