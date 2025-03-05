package postgres

import (
	"fmt"
	"log"

	"github.com/gene-qxsi/Blog/user-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	const op = "user-service>internal>infrastructure>postgres>postgres.go>NewPostgresDB()"
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error: %s. op: %s", err.Error(), op)
	}

	return db, nil
}
