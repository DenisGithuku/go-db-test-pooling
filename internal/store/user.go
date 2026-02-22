package store

import (
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	ID int
	Name string
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetById(ctx context.Context, id int) (*User, error)
	Save(ctx context.Context, u *User) (*User, error)
}

type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRepository (pool *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{pool: pool}
}

func (r *PostgresUserRepository) Save(ctx context.Context, u *User) (*User, error) {
	err := r.pool.QueryRow(ctx, `INSERT INTO users (name) VALUES ($1) RETURNING id`, u.Name).Scan(&u.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
	
}

func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]*User, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, name FROM users`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresUserRepository) GetById(ctx context.Context, id int) (*User, error) {
	var u User
	err := r.pool.QueryRow(ctx, `SELECT id, name FROM users WHERE id=$1`, id).Scan(&u.ID, &u.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}