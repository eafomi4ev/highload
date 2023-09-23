package db_pg

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"otus_highload/internal/domain"
)

type PostgresDB struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *PostgresDB {
	return &PostgresDB{
		conn: conn,
	}
}

// todo: заменить тип user_register.User на тип из текущего слоя, чтобы не было протечки абстракции
func (d *PostgresDB) AddUser(ctx context.Context, user domain.User) (domain.User, error) {
	var id string
	err := d.conn.QueryRow(
		ctx,
		"insert into users(id, first_name, surname, birthdate, biography, password_hash, city_id) values ($1, $2, $3, $4, $5, $6, $7) returning id",
		user.ID,
		user.FirstName,
		user.Surname,
		user.Birthdate,
		user.Biography,
		user.PasswordHash,
		user.City.ID,
	).Scan(&id)
	if err != nil {
		return domain.User{}, err
	}

	user.ID = uuid.MustParse(id)

	return user, nil
}

func (d *PostgresDB) GetCityByName(ctx context.Context, name string) (domain.City, error) {
	var city domain.City
	err := d.conn.QueryRow(
		ctx,
		"select id, name from cities where name=$1",
		name,
	).Scan(&city.ID, &city.Name)
	if err != nil {
		return domain.City{}, err
	}

	return city, nil
}
