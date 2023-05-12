package db

import (
	"database/sql"
	"fmt"
	"github.com/KutsDenis/anti-smoking/internal/config"
	"time"
)

func NewDB() (*sql.DB, error) {
	cfg := config.Get.DB

	connStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 3; i++ {
		if err = db.Ping(); err != nil {
			if i == 2 {
				return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
			}
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	return db, nil
}
