package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const maskText = "***"

// PostgresConfig конфигурация БД
type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	DBName   string `mapstructure:"dbname"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
}

// DB represents the database connection
type DB struct {
	Pool *pgxpool.Pool
}

type Pg struct {
	dbURL string
	Db    *DB
}

// NewPostgres получить новый интерфейс для работы с БД
func NewPostgres(dbURL string) *Pg {
	return &Pg{dbURL: dbURL}
}

// InitPg инициализация соединения
func (p *Pg) InitPg() error {
	db, err := p.New(p.dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Connected to database successfully")

	p.Db = db

	return nil
}

// MaskText маскирует строку заменяя середину на "***"
func (p *Pg) MaskText(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		s = s + s[:1]
	}

	return s[:2] + maskText + s[len(s)-2:]
}

// New creates a new DB instance
func (p *Pg) New(connString string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &DB{Pool: pool}, nil
}

// Close closes the database connection
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
