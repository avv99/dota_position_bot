package PostgreSQL

import (
	"database/sql"
	"dota_position_bot/internal/storage"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func InitPostgresStorage(connStr string) (storage.Storage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	pg := &Postgres{db: db}
	flag, err := pg.TableExists("heroes")
	if err != nil {
		return nil, err
	}
	if !flag {
		err = pg.CreateTable()
		if err != nil {
			return nil, err
		}
		err = pg.LoadData()
		if err != nil {
			return nil, err
		}
	}
	return pg, nil
}

func (p *Postgres) GetHeroes(pos string) ([]string, error) {
	query := "SELECT hero FROM heroes WHERE position = $1"
	rows, err := p.db.Query(query, pos)
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		return []string{}, err
	}
	defer rows.Close()

	var heroes []string
	for rows.Next() {
		var hero string
		if err := rows.Scan(&hero); err != nil {
			log.Println("Ошибка при чтении результата запроса:", err)
			return []string{}, err
		}
		heroes = append(heroes, hero)
	}

	if len(heroes) == 0 {
		log.Println("Выбрана неправильная позиция")
		return []string{}, errors.New("Выбрана неправильная позиция")
	}

	return heroes, nil
}

func (p *Postgres) LoadData() error {
	data := map[string][]string{
		"1pos": {"Windranger", "Broodmother", "Clinkz", "MonkeyKing"},
		"2pos": {"Tiny", "Broodmother", "EmberSpirit", "PrimalBeast"},
		"3pos": {"DarkSeer", "Broodmother", "Enigma", "Windranger"},
		"4pos": {"SpiritBreaker", "Clockwerk", "ShadowDemon", "Weawer"},
		"5pos": {"Sven", "ElderTitan", "TreantProtector", "Weawer"},
	}

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	for pos, heroes := range data {
		for _, hero := range heroes {
			_, err := tx.Exec("INSERT INTO heroes (position, hero) VALUES ($1, $2)", pos, hero)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) TableExists(tableName string) (bool, error) {
	query := "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = $1)"
	var exists bool
	err := p.db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (p *Postgres) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS heroes (
		id SERIAL PRIMARY KEY,
		position VARCHAR(255),
		hero VARCHAR(255)
	)`
	_, err := p.db.Exec(query)
	if err != nil {
		log.Println("Ошибка при создании таблицы:", err)
		return err
	}

	return nil
}
