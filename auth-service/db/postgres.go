package db

import (
	models "auth-service/model"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func InitPostgres() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("POSTGRES_DSN is not set")
	}

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Postgres connection failed:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Postgres ping failed:", err)
	}
}

//func InitPostgres() {
//	dsn := os.Getenv("postgresql://postgres:pupipu@localhost/postgres?client_encoding=utf8") // пример: postgres://user:pass@user-db:5432/dbname?sslmode=disable
//	var err error
//	DB, err = sql.Open("postgres", dsn)
//	if err != nil {
//		log.Fatal("Postgres connection failed:", err)
//	}
//}

func SaveUser(u models.User) error {
	_, err := DB.Exec("INSERT INTO users_info (email, login, password) VALUES ($1, $2, $3)", u.Email, u.Login, u.Password)
	return err
}

func GetUserByToken(token string) (*models.User, error) {
	row := DB.QueryRow("SELECT id, email, login, password FROM users_info WHERE login = $1", token)
	var u models.User
	err := row.Scan(&u.ID, &u.Email, &u.Login, &u.Password)
	return &u, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetUserByLogin(login string) (*models.User, error) {
	row := DB.QueryRow("SELECT id, email, login, password FROM users_info WHERE login = $1", login)
	var u models.User
	err := row.Scan(&u.ID, &u.Email, &u.Login, &u.Password)
	return &u, err
}
